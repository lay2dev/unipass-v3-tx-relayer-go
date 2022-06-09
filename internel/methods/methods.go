package methods

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gin-gonic/gin"

	"asset_forwarder/internel/contracts/asset"
	"asset_forwarder/internel/types"
	"asset_forwarder/internel/utils"
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
		nonce, err = h.nonceCache.getNonceAndAdd()
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
	assetTxJson := &types.TransferNativeTxJson{}
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

	if amount, ok := h.FeeTokens[assetTx.FeeToken]; !ok || assetTx.FeeAmount.Cmp(amount) < 0 {
		c.IndentedJSON(http.StatusInternalServerError, "fee amount is not enough")
	}

	asset, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	auth, err := h.genAuth()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	tx, err := asset.TransferNativeToken(auth,
		assetTx.Nonce,
		assetTx.To,
		assetTx.Amount,
		assetTx.FeeToken,
		assetTx.FeeAmount,
		assetTx.Sig,
		assetTx.SigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	h.txList.AddTx(tx)
	log.Info("TransferNative tx sent, hash:", tx.Hash())
	c.IndentedJSON(http.StatusOK, tx.Hash())
}

func (h *AssetTxHandler) TransferToken(c *gin.Context) {
	assetTxJson := &types.TransferTokenTxJson{}
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

	if amount, ok := h.FeeTokens[assetTx.FeeToken]; !ok || assetTx.FeeAmount.Cmp(amount) < 0 {
		c.IndentedJSON(http.StatusInternalServerError, "fee amount is not enough")
	}

	asset, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	auth, err := h.genAuth()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	tx, err := asset.TransferToken(auth,
		assetTx.Nonce,
		assetTx.Token,
		assetTx.To,
		assetTx.Amount,
		assetTx.FeeToken,
		assetTx.FeeAmount,
		assetTx.Sig,
		assetTx.SigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	h.txList.AddTx(tx)

	log.Info("TransferToken tx sent, hash:", tx.Hash())
	c.IndentedJSON(http.StatusOK, tx.Hash())
}

func (h *AssetTxHandler) Execute(c *gin.Context) {
	assetTxJson := &types.ExecuteTxJson{}
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

	if amount, ok := h.FeeTokens[assetTx.FeeToken]; !ok || assetTx.FeeAmount.Cmp(amount) < 0 {
		c.IndentedJSON(http.StatusInternalServerError, "fee amount is not enough")
	}

	asset, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	auth, err := h.genAuth()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	tx, err := asset.Execute(auth,
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

	h.txList.AddTx(tx)
	log.Info("Execute Tx sent, hash:", tx.Hash())
	c.IndentedJSON(http.StatusOK, tx.Hash())
}
