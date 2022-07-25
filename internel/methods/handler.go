package methods

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"

	"tx_relayer/internel/configs"
	"tx_relayer/internel/contracts/asset"
	"tx_relayer/internel/contracts/entry"
	"tx_relayer/internel/logger"
	"tx_relayer/internel/utils"
)

var log = logger.SugarLogger

type AssetTxHandler struct {
	client        *ethclient.Client
	webhookClient webhook.Client
	chainID       *big.Int
	entry         *entry.Entry
	assetABI      abi.ABI
	privKey       *ecdsa.PrivateKey
	txList        *TransactionList
	discordId     snowflake.ID
	discordToken  string
	feeTokens     map[common.Address]*big.Int
}

const API_URL = "api_url"
const ENTRY_ADDRESS = "entry_address"
const FEE_PROVIDER = "fee_provider"

func NewAssetTxHandler(conf *configs.ForwarderConfig) (*AssetTxHandler, error) {
	h := &AssetTxHandler{
		feeTokens: make(map[common.Address]*big.Int),
	}

	for addrStr, amount := range conf.FeeTokens {
		addr := common.HexToAddress(addrStr)
		h.feeTokens[addr] = new(big.Int).SetInt64(amount)
	}

	client, err := ethclient.Dial(conf.ApiUrl)
	if err != nil {
		return nil, err
	}

	var entryAddress common.Address
	{
		buf, err := utils.HexToBytes(conf.EntryAddress)
		if err != nil {
			return nil, err
		}
		copy(entryAddress[:], buf)
	}
	entry, err := entry.NewEntry(entryAddress, client)
	if err != nil {
		return nil, err
	}

	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(conf.FeeProvider, "0x"))
	if err != nil {
		return nil, err
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	currentNonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	txList := NewTransactionList(fromAddress, currentNonce)

	h.assetABI, err = abi.JSON(strings.NewReader(asset.AssetABI))
	if err != nil {
		return nil, err
	}
	h.discordId = snowflake.ID(conf.DiscordId)
	h.discordToken = conf.DiscordToken
	h.client = client
	h.chainID = chainID
	h.entry = entry
	h.privKey = privKey
	h.txList = txList

	if h.discordId != 0 && h.discordToken != "" {
		h.webhookClient = webhook.New(h.discordId, h.discordToken)
		fmt.Printf("relayer %s started", h.txList.fromAddress)
		sendDiscord(h.webhookClient, fmt.Sprintf("relayer %s started", h.txList.fromAddress))
	}

	fmt.Println("Entry address: ", entryAddress)
	fmt.Println("fee provider address: ", fromAddress)
	fmt.Println("chainID: ", chainID)

	go h.Monitor()
	go func() {
		for {
			h.Run()
		}
	}()

	return h, nil
}

func (h *AssetTxHandler) Run() {
	defer func() {
		if err := recover(); err != nil {
			// 打印异常，关闭资源，退出此函数
			log.Errorf("relayer paniced, err: %s", err)
			if h.webhookClient != nil {
				sendDiscord(h.webhookClient, fmt.Sprintf("relayer %s paniced, begin to restart...", h.txList.fromAddress))
			}
		}
	}()

	count := int64(10)
	for {
		if txs := h.txList.GetTx(count); txs != nil {
			log.Infof("begin to send %d transactions to blockchain, current txlist minnumber:%d, maxnumber: %d", len(txs), h.txList.minNumber, h.txList.maxNumber)
			for _, tx := range txs {
				h.SendTransactionAndWait(tx)
			}
			finishedCount := len(txs)
			h.txList.FinishTx(int64(finishedCount))
			log.Infof("send %d transactions to blockchain finished, current txlist minnumber:%d, maxnumber: %d", len(txs), h.txList.minNumber, h.txList.maxNumber)
			continue
		}
		time.Sleep(time.Second)
	}
}

func (h *AssetTxHandler) Monitor() {
	ticker := time.NewTicker(120 * time.Second)
	lastNonce := uint64(0)

	for range ticker.C {
		tx := h.txList.txs[h.txList.minNumber]
		if tx != nil {
			if lastNonce == tx.Nonce() {
				if h.webhookClient != nil {
					sendDiscord(h.webhookClient, fmt.Sprintf("relayer %s get stuck", h.txList.fromAddress))
				} else {
					log.Error("relayer get stuck")
				}
			}
			lastNonce = tx.Nonce()
		} else {
			ok := h.txList.TryLock()
			if !ok {
				if h.webhookClient != nil {
					sendDiscord(h.webhookClient, fmt.Sprintf("relayer %s maybe deadlocked", h.txList.fromAddress))
				} else {
					log.Error("relayer get stuck")
				}
			} else {
				h.txList.Unlock()
			}
		}
	}
}

// send(s) a message to the webhook
func sendDiscord(client webhook.Client, mes string) {
	if _, err := client.CreateMessage(discord.NewWebhookMessageCreateBuilder().
		SetContentf(mes).
		Build(),
	); err != nil {
		log.Errorf("error sending message:{%s} to discord %s", mes, err)
	}
}

func (h *AssetTxHandler) SendTransactionAndWait(tx *types.Transaction) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	for {
		err := h.client.SendTransaction(ctx, tx)
		if err != nil {
			// 判断交易是否已经上链了
			receipt, err := h.client.TransactionReceipt(ctx, tx.Hash())
			if err == nil {
				if receipt.Status == 1 {
					log.Infof("Tx: %s execute success", tx.Hash())
				} else {
					log.Infof("Tx: %s execute failed", tx.Hash())
				}
				return
			}
			// 否则继续尝试发送
			log.Error("SendTransaction:", tx.Hash(), " failed, retrying.....")
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	log.Infof("Send Tx: %s to node success, waiting finality", tx.Hash())
	// 等100 ms
	time.Sleep(100 * time.Millisecond)
	retryCount := 0
	for {
		receipt, err := h.client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			if receipt.Status == 1 {
				log.Infof("Tx: %s execute success", tx.Hash())
			} else {
				log.Infof("Tx: %s execute failed", tx.Hash())
			}
			return
		}
		if err == ethereum.NotFound {
			if retryCount < 5 {
				log.Error("TransactionReceipt:", tx.Hash(), " not found, retrying.....")
				retryCount++
				time.Sleep(2 * time.Second)
				continue
			}
		}
		log.Errorf("Tx: %s get receipt error: %s", tx.Hash(), err)
		return
	}
}
