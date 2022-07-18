package methods

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/gin-gonic/gin"

	"tx_relayer/internel/contracts/asset"
	local_types "tx_relayer/internel/types"
	"tx_relayer/internel/utils"
)

var (
	errTransactionTooMuch = errors.New("transaction too much")
)

func (h *AssetTxHandler) genAuth() (*bind.TransactOpts, error) {
	gasPrice, err := h.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(h.privKey, h.chainID)
	if err != nil {
		return nil, err
	}

	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice
	var nonce uint64

	retryCount := 0
	for {
		nonce, err = h.txList.getNonceAndAdd()
		if err == nil {
			break
		}
		if err == errTransactionTooMuch && retryCount < 3 {
			retryCount++
			time.Sleep(time.Second * 2)
			continue
		}
		if err != nil {
			return nil, err
		}
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.NoSend = true

	return auth, nil
}

func (h *AssetTxHandler) GetChainId(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.chainID)
}

func (h *AssetTxHandler) GetAssetAddress(c *gin.Context) {
	registerEmailStr := c.Param("register_email")
	registerEmail, err := utils.HexToBytes32(registerEmailStr)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.entry.Users(&bind.CallOpts{}, registerEmail)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res.AssetContractAddress)
}

func (h *AssetTxHandler) TransferNative(c *gin.Context) {
	assetTxJson := &local_types.TransferNativeTxJson{}
	err := c.BindJSON(assetTxJson)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	{
		res, _ := json.Marshal(assetTxJson)
		log.Info("Handler TransferNative Tx: ", string(res))
	}

	assetTx, err := assetTxJson.Convert()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	if amount, ok := h.feeTokens[assetTx.FeeToken]; !ok || assetTx.FeeAmount.Cmp(amount) < 0 {
		c.IndentedJSON(http.StatusInternalServerError, "fee amount is not enough")
	}

	assetContract, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	gasPrice, err := h.client.SuggestGasPrice(context.Background())
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	callData, err := h.assetABI.Pack("transferNativeToken",
		assetTx.Nonce,
		assetTx.To,
		assetTx.Amount,
		assetTx.FeeToken,
		assetTx.FeeAmount,
		assetTx.Sig,
		assetTx.SigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	_, err = h.client.CallContract(c, ethereum.CallMsg{
		From:      h.txList.fromAddress,
		To:        &assetTx.AssetContract,
		Gas:       1000000,
		GasPrice:  gasPrice,
		GasFeeCap: gasPrice,
		GasTipCap: gasPrice,
		Value:     big.NewInt(0),
		Data:      callData,
	}, nil)

	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	tx, err := h.addTransferNativeTx(assetContract, assetTx)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	c.IndentedJSON(http.StatusOK, tx.Hash())
}

func (h *AssetTxHandler) addTransferNativeTx(
	assetContract *asset.Asset,
	assetTx *local_types.TransferNativeTx) (*types.Transaction, error) {

	h.txList.Lock()
	defer h.txList.Unlock()

	auth, err := h.genAuth()
	if err != nil {
		return nil, err
	}

	tx, err := assetContract.TransferNativeToken(auth,
		assetTx.Nonce,
		assetTx.To,
		assetTx.Amount,
		assetTx.FeeToken,
		assetTx.FeeAmount,
		assetTx.Sig,
		assetTx.SigKeyIndex)
	if err != nil {
		return nil, err
	}

	h.txList.AddTx(tx)
	log.Infof("TransferNative tx added, hash:%s, nonce: %d", tx.Hash(), tx.Nonce())
	return tx, nil
}

func (h *AssetTxHandler) TransferToken(c *gin.Context) {
	assetTxJson := &local_types.TransferTokenTxJson{}
	err := c.BindJSON(assetTxJson)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	{
		res, _ := json.Marshal(assetTxJson)
		log.Info("Handler TransferToken Tx: ", string(res))
	}

	assetTx, err := assetTxJson.Convert()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	if amount, ok := h.feeTokens[assetTx.FeeToken]; !ok || assetTx.FeeAmount.Cmp(amount) < 0 {
		c.IndentedJSON(http.StatusInternalServerError, "fee amount is not enough")
	}

	assetContract, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	gasPrice, err := h.client.SuggestGasPrice(context.Background())
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	callData, err := h.assetABI.Pack("transferToken",
		assetTx.Nonce,
		assetTx.Token,
		assetTx.To,
		assetTx.Amount,
		assetTx.FeeToken,
		assetTx.FeeAmount,
		assetTx.Sig,
		assetTx.SigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	_, err = h.client.CallContract(c, ethereum.CallMsg{
		From:      h.txList.fromAddress,
		To:        &assetTx.AssetContract,
		Gas:       1000000,
		GasPrice:  gasPrice,
		GasFeeCap: gasPrice,
		GasTipCap: gasPrice,
		Value:     big.NewInt(0),
		Data:      callData,
	}, nil)

	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	tx, err := h.addTransferTokenTx(assetContract, assetTx)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	c.IndentedJSON(http.StatusOK, tx.Hash())
}

func (h *AssetTxHandler) addTransferTokenTx(
	assetContract *asset.Asset,
	assetTx *local_types.TransferTokenTx) (*types.Transaction, error) {
	h.txList.Lock()
	defer h.txList.Unlock()

	auth, err := h.genAuth()
	if err != nil {
		return nil, err
	}

	tx, err := assetContract.TransferToken(auth,
		assetTx.Nonce,
		assetTx.Token,
		assetTx.To,
		assetTx.Amount,
		assetTx.FeeToken,
		assetTx.FeeAmount,
		assetTx.Sig,
		assetTx.SigKeyIndex)
	if err != nil {
		return nil, err
	}

	h.txList.AddTx(tx)

	log.Infof("TransferToken tx added, hash:%s, nonce:%d", tx.Hash(), tx.Nonce())
	return tx, nil
}

func (h *AssetTxHandler) Execute(c *gin.Context) {
	assetTxJson := &local_types.ExecuteTxJson{}
	err := c.BindJSON(assetTxJson)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	{
		res, _ := json.Marshal(assetTxJson)
		log.Info("Handler Execute Tx: ", string(res))
	}

	assetTx, err := assetTxJson.Convert()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	if amount, ok := h.feeTokens[assetTx.FeeToken]; !ok || assetTx.FeeAmount.Cmp(amount) < 0 {
		c.IndentedJSON(http.StatusInternalServerError, "fee amount is not enough")
	}

	assetContract, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	gasPrice, err := h.client.SuggestGasPrice(context.Background())
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	callData, err := h.assetABI.Pack("execute",
		assetTx.Nonce,
		assetTx.To,
		assetTx.Value,
		assetTx.CallData,
		assetTx.FeeToken,
		assetTx.FeeAmount,
		assetTx.Sig,
		assetTx.SigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	_, err = h.client.CallContract(c, ethereum.CallMsg{
		From:      h.txList.fromAddress,
		To:        &assetTx.AssetContract,
		Gas:       1000000,
		GasPrice:  gasPrice,
		GasFeeCap: gasPrice,
		GasTipCap: gasPrice,
		Value:     big.NewInt(0),
		Data:      callData,
	}, nil)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	tx, err := h.addExecuteTx(assetContract, assetTx)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	c.IndentedJSON(http.StatusOK, tx.Hash())
}

func (h *AssetTxHandler) addExecuteTx(
	assetContract *asset.Asset,
	assetTx *local_types.ExecuteTx) (*types.Transaction, error) {
	h.txList.Lock()
	defer h.txList.Unlock()

	auth, err := h.genAuth()
	if err != nil {
		return nil, err
	}
	tx, err := assetContract.Execute(auth,
		assetTx.Nonce,
		assetTx.To,
		assetTx.Value,
		assetTx.CallData,
		assetTx.FeeToken,
		assetTx.FeeAmount,
		assetTx.Sig,
		assetTx.SigKeyIndex)
	if err != nil {
		return nil, err
	}

	h.txList.AddTx(tx)

	log.Infof("Execute Tx added, hash:%s, nonce:%d", tx.Hash(), tx.Nonce())
	return tx, nil
}
