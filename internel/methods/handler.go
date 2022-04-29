package methods

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"asset_forwarder/internel/configs"
	"asset_forwarder/internel/contracts/entry"
	"asset_forwarder/internel/utils"
)

type AssetTxHandler struct {
	client      *ethclient.Client
	chainID     *big.Int
	entry       *entry.Entry
	privKey     *ecdsa.PrivateKey
	fromAddress common.Address
}

const API_URL = "api_url"
const ENTRY_ADDRESS = "entry_address"
const FEE_PROVIDER = "fee_provider"

func NewAssetTxHandler() (*AssetTxHandler, error) {
	viper, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	apiUrl := viper.GetString(API_URL)
	client, err := ethclient.Dial(apiUrl)
	if err != nil {
		return nil, err
	}
	entryAddressStr := viper.GetString(ENTRY_ADDRESS)

	var entryAddress common.Address
	{
		buf, err := utils.HexToBytes(entryAddressStr)
		if err != nil {
			panic(err)
		}
		copy(entryAddress[:], buf)
	}
	entry, err := entry.NewEntry(entryAddress, client)
	if err != nil {
		panic(err)
	}

	feeProvider := viper.GetString(FEE_PROVIDER)
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(feeProvider, "0x"))
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
	return &AssetTxHandler{
		client,
		chainID,
		entry,
		privKey,
		fromAddress,
	}, nil
}
