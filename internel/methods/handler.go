package methods

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"tx_relayer/internel/configs"
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
			panic(err)
		}
		copy(entryAddress[:], buf)
	}
	entry, err := entry.NewEntry(entryAddress, client)
	if err != nil {
		panic(err)
	}

	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(conf.FeeProvider, "0x"))
	if err != nil {
		panic(err)
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
	for {
		if tx := h.txList.GetTx(); tx != nil {
			h.SendTransactionAndWait(tx)
			h.txList.FinishTx()
			continue
		}
		time.Sleep(time.Second)
	}
}

func (h *AssetTxHandler) SendTransactionAndWait(tx *types.Transaction) {
	ctx := context.Background()
	err := h.client.SendTransaction(ctx, tx)
	if err != nil {
		return
	}
	retryCount := 0
	for {
		receipt, err := h.client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			if receipt.Status == 1 {
				return
			}
			return
		}
		if err == ethereum.NotFound && retryCount < 3 {
			retryCount++
			time.Sleep(3 * time.Second)
			continue
		}
		log.Infof("TransferNative tx sent, hash:%s, err:%v\n", tx.Hash(), err)
		return
	}
}
