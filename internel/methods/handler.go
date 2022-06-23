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

	"tx_relayer/internel/configs"
	"tx_relayer/internel/contracts/asset"
	"tx_relayer/internel/contracts/entry"
	"tx_relayer/internel/logger"
	"tx_relayer/internel/utils"
)

var log = logger.SugarLogger

type AssetTxHandler struct {
	client     *ethclient.Client
	nonceCache *NonceCache
	chainID    *big.Int
	entry      *entry.Entry
	assetABI   abi.ABI
	privKey    *ecdsa.PrivateKey
	txList     *TransactionList
	feeTokens  map[common.Address]*big.Int
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
	nonceCache := newNonceCache(fromAddress, currentNonce)

	txList := NewTransactionList()

	h.assetABI, err = abi.JSON(strings.NewReader(asset.AssetABI))
	if err != nil {
		return nil, err
	}
	h.client = client
	h.nonceCache = nonceCache
	h.chainID = chainID
	h.entry = entry
	h.privKey = privKey
	h.txList = txList

	fmt.Println("Entry address: ", entryAddress)
	fmt.Println("fee provider address: ", fromAddress)
	fmt.Println("chainID: ", chainID)

	go h.Run()

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
