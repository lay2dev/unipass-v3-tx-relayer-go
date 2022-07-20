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
	client       *ethclient.Client
	chainID      *big.Int
	entry        *entry.Entry
	assetABI     abi.ABI
	privKey      *ecdsa.PrivateKey
	txList       *TransactionList
	discordId    snowflake.ID
	discordToken string
	feeTokens    map[common.Address]*big.Int
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

	fmt.Println("Entry address: ", entryAddress)
	fmt.Println("fee provider address: ", fromAddress)
	fmt.Println("chainID: ", chainID)

	go h.Run()
	go h.Monitor()

	return h, nil
}

func (h *AssetTxHandler) Run() {
	count := int64(10)
	for {
		if txs := h.txList.GetTx(count); txs != nil {
			for _, tx := range txs {
				h.SendTransactionAndWait(tx)
			}
			finishedCount := len(txs)
			h.txList.FinishTx(int64(finishedCount))
			continue
		}
		time.Sleep(time.Second)
	}
}

func (h *AssetTxHandler) Monitor() {
	ticker := time.NewTicker(120 * time.Second)
	lastNonce := uint64(0)
	var client webhook.Client
	if h.discordId != 0 && h.discordToken != "" {
		client = webhook.New(h.discordId, h.discordToken)
		fmt.Printf("relayer %s started", h.txList.fromAddress)
		sendDiscord(client, fmt.Sprintf("relayer %s started", h.txList.fromAddress))
	}

	for range ticker.C {
		tx := h.txList.txs[h.txList.minNumber]
		if tx != nil {
			if lastNonce == tx.Nonce() {
				if client != nil {
					sendDiscord(client, fmt.Sprintf("relayer %s get stuck", h.txList.fromAddress))
				} else {
					log.Error("relayer get stuck")
				}
			}
			lastNonce = tx.Nonce()
		} else {
			ok := h.txList.TryLock()
			if !ok {
				if client != nil {
					sendDiscord(client, fmt.Sprintf("relayer %s maybe deadlocked", h.txList.fromAddress))
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
	ctx := context.Background()
	for {
		err := h.client.SendTransaction(ctx, tx)
		if err != nil {
			log.Error("SendTransaction:", tx, " failed, retrying.....")
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
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
				retryCount++
				time.Sleep(3 * time.Second)
				continue
			}
		}
		log.Errorf("Tx: %s get receipt errir", tx.Hash())
		break
	}
}
