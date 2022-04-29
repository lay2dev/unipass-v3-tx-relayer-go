package types

import (
	"asset_forwarder/internel/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TransferNativeTxJson struct {
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

type TransferNativeTx struct {
	AssetContract common.Address
	Nonce         *big.Int
	To            common.Address
	Amount        *big.Int
	FeeToken      common.Address
	FeeAmount     *big.Int
	Sig           []byte
	SigKeyIndex   *big.Int
}

func (tx *TransferNativeTxJson) Convert() (*TransferNativeTx, error) {
	nonce, err := utils.BigIntDecode(tx.Nonce)
	if err != nil {
		return nil, err
	}
	amount, err := utils.BigIntDecode(tx.Amount)
	if err != nil {
		return nil, err
	}
	feeAmount, err := utils.BigIntDecode(tx.FeeAmount)
	if err != nil {
		return nil, err
	}
	sig, err := utils.HexToBytes(tx.Sig)
	if err != nil {
		return nil, err
	}
	sigKeyIndex, err := utils.BigIntDecode(tx.SigKeyIndex)
	if err != nil {
		return nil, err
	}

	return &TransferNativeTx{
		AssetContract: tx.AssetContract,
		Nonce:         nonce,
		To:            tx.To,
		Amount:        amount,
		FeeToken:      tx.FeeToken,
		FeeAmount:     feeAmount,
		Sig:           sig,
		SigKeyIndex:   sigKeyIndex,
	}, nil
}

type TransferTokenTxJson struct {
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

type TransferTokenTx struct {
	AssetContract common.Address
	Nonce         *big.Int
	Token         common.Address
	To            common.Address
	Amount        *big.Int
	FeeToken      common.Address
	FeeAmount     *big.Int
	Sig           []byte
	SigKeyIndex   *big.Int
}

func (tx *TransferTokenTxJson) Convert() (*TransferTokenTx, error) {
	nonce, err := utils.BigIntDecode(tx.Nonce)
	if err != nil {
		return nil, err
	}
	amount, err := utils.BigIntDecode(tx.Amount)
	if err != nil {
		return nil, err
	}
	feeAmount, err := utils.BigIntDecode(tx.FeeAmount)
	if err != nil {
		return nil, err
	}
	sig, err := utils.HexToBytes(tx.Sig)
	if err != nil {
		return nil, err
	}
	sigKeyIndex, err := utils.BigIntDecode(tx.SigKeyIndex)
	if err != nil {
		return nil, err
	}

	return &TransferTokenTx{
		AssetContract: tx.AssetContract,
		Nonce:         nonce,
		Token:         tx.Token,
		To:            tx.To,
		Amount:        amount,
		FeeToken:      tx.FeeToken,
		FeeAmount:     feeAmount,
		Sig:           sig,
		SigKeyIndex:   sigKeyIndex,
	}, nil
}

type ExecuteTxJson struct {
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

type ExecuteTx struct {
	AssetContract common.Address
	Nonce         *big.Int
	To            common.Address
	Value         *big.Int
	CallData      []byte
	FeeToken      common.Address
	FeeAmount     *big.Int
	Sig           []byte
	SigKeyIndex   *big.Int
}

func (tx *ExecuteTxJson) Convert() (*ExecuteTx, error) {
	nonce, err := utils.BigIntDecode(tx.Nonce)
	if err != nil {
		return nil, err
	}
	value, err := utils.BigIntDecode(tx.Value)
	if err != nil {
		return nil, err
	}
	callData, err := utils.HexToBytes(tx.CallData)
	if err != nil {
		return nil, err
	}
	feeAmount, err := utils.BigIntDecode(tx.FeeAmount)
	if err != nil {
		return nil, err
	}
	sig, err := utils.HexToBytes(tx.Sig)
	if err != nil {
		return nil, err
	}
	sigKeyIndex, err := utils.BigIntDecode(tx.SigKeyIndex)
	if err != nil {
		return nil, err
	}

	return &ExecuteTx{
		AssetContract: tx.AssetContract,
		Nonce:         nonce,
		To:            tx.To,
		Value:         value,
		CallData:      callData,
		FeeToken:      tx.FeeToken,
		FeeAmount:     feeAmount,
		Sig:           sig,
		SigKeyIndex:   sigKeyIndex,
	}, nil
}
