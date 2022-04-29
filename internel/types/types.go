package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type TransferNativeTx struct {
	AssetContract common.Address `json:"assetContract"`
	RegisterEmail string         `json:"registerEmail"`
	Nonce         string         `json:"nonce"`
	To            common.Address `json:"to"`
	Amount        string         `json:"amount"`
	FeeToken      common.Address `json:"feeToken"`
	FeeAmount     string         `json:"feeAmount"`
	Sig           string         `json:"sig"`
	SigKeyIndex   string         `json:"sigKeyIndex"`
}

type TransferTokenTx struct {
	AssetContract common.Address `json:"assetContract"`
	RegisterEmail string         `json:"registerEmail"`
	Nonce         string         `json:"nonce"`
	Token         common.Address `json:"token"`
	To            common.Address `json:"to"`
	Amount        string         `json:"amount"`
	FeeToken      common.Address `json:"feeToken"`
	FeeAmount     string         `json:"feeAmount"`
	Sig           string         `json:"sig"`
	SigKeyIndex   string         `json:"sigKeyIndex"`
}

type ExecuteTx struct {
	AssetContract common.Address `json:"assetContract"`
	RegisterEmail string         `json:"registerEmail"`
	Nonce         string         `json:"nonce"`
	To            common.Address `json:"to"`
	Value         string         `json:"value"`
	CallData      string         `json:"callData"`
	FeeToken      common.Address `json:"feeToken"`
	FeeAmount     string         `json:"feeAmount"`
	Sig           string         `json:"sig"`
	SigKeyIndex   string         `json:"sigKeyIndex"`
}
