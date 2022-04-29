package methods

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gin-gonic/gin"

	"asset_forwarder/internel/contracts/asset"
	"asset_forwarder/internel/logger"
	"asset_forwarder/internel/types"
	"asset_forwarder/internel/utils"
)

func (h *AssetTxHandler) genAuth() (*bind.TransactOpts, error) {
	nonce, err := h.client.PendingNonceAt(context.Background(), h.fromAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := h.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(h.privKey, h.chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice
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
		logger.Info("Handler TransferNative Tx: ", string(res))
	}

	assetTx, err := assetTxJson.Convert()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
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
	retryCount := 0
	for {
		receipt, err := h.client.TransactionReceipt(c, tx.Hash())
		if err == nil {
			if receipt.Status == 1 {
				c.IndentedJSON(http.StatusOK, tx.Hash())
				return
			}
			c.IndentedJSON(http.StatusBadRequest, tx.Hash())
			return
		}
		if err == ethereum.NotFound && retryCount < 3 {
			retryCount++
			time.Sleep(3 * time.Second)
			continue
		}
		c.IndentedJSON(http.StatusBadRequest, tx.Hash())
		return
	}
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
		logger.Info("Handler TransferToken Tx: ", string(res))
	}

	assetTx, err := assetTxJson.Convert()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
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
	retryCount := 0
	for {
		receipt, err := h.client.TransactionReceipt(c, tx.Hash())
		if err == nil {
			if receipt.Status == 1 {
				c.IndentedJSON(http.StatusOK, tx.Hash())
				return
			}
			c.IndentedJSON(http.StatusBadRequest, tx.Hash())
			return
		}
		if err == ethereum.NotFound && retryCount < 3 {
			retryCount++
			time.Sleep(3 * time.Second)
			continue
		}
		c.IndentedJSON(http.StatusBadRequest, tx.Hash())
		return
	}
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
		logger.Info("Handler Execute Tx: ", string(res))
	}

	assetTx, err := assetTxJson.Convert()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
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
	{
		logger.Info("Execute Tx sent, hash:", tx.Hash())
	}
	retryCount := 0
	for {
		receipt, err := h.client.TransactionReceipt(c, tx.Hash())
		if err == nil {
			if receipt.Status == 0 {
				c.IndentedJSON(http.StatusBadRequest, tx.Hash())
				return
			}
			for _, log := range receipt.Logs {
				_, err := asset.ParseECallSuccess(*log)
				if err == nil {
					c.IndentedJSON(http.StatusOK, tx.Hash())
					return
				}
			}

			c.IndentedJSON(http.StatusBadRequest, tx.Hash())
			return
		}
		if err == ethereum.NotFound && retryCount < 3 {
			retryCount++
			time.Sleep(3 * time.Second)
			continue
		}
		c.IndentedJSON(http.StatusBadRequest, tx.Hash())
		return
	}
}
