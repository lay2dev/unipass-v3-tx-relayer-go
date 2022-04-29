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
	assetTx := &types.TransferNativeTx{}
	err := c.BindJSON(assetTx)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	{
		res, _ := json.Marshal(assetTx)
		logger.Info("Handler TransferToken Tx: ", string(res))
	}
	asset, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	auth, err := h.genAuth()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	nonce, err := utils.BigIntDecode(assetTx.Nonce)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	amount, err := utils.BigIntDecode(assetTx.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	feeAmount, err := utils.BigIntDecode(assetTx.FeeAmount)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	sig, err := utils.HexToBytes(assetTx.Sig)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	sigKeyIndex, err := utils.BigIntDecode(assetTx.SigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	tx, err := asset.TransferNativeToken(auth, nonce, assetTx.To, amount, assetTx.FeeToken, feeAmount, sig, sigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
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
	assetTx := &types.TransferTokenTx{}
	err := c.BindJSON(assetTx)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	{
		res, _ := json.Marshal(assetTx)
		logger.Info("Handler TransferToken Tx: ", string(res))
	}
	asset, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	auth, err := h.genAuth()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	nonce, err := utils.BigIntDecode(assetTx.Nonce)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	amount, err := utils.BigIntDecode(assetTx.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	feeAmount, err := utils.BigIntDecode(assetTx.FeeAmount)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	sig, err := utils.HexToBytes(assetTx.Sig)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	sigKeyIndex, err := utils.BigIntDecode(assetTx.SigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	tx, err := asset.TransferToken(auth, nonce, assetTx.Token, assetTx.To, amount, assetTx.FeeToken, feeAmount, sig, sigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
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
	assetTx := &types.ExecuteTx{}
	err := c.BindJSON(assetTx)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	{
		res, _ := json.Marshal(assetTx)
		logger.Info("Handler TransferToken Tx: ", string(res))
	}
	asset, err := asset.NewAsset(assetTx.AssetContract, h.client)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	auth, err := h.genAuth()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	nonce, err := utils.BigIntDecode(assetTx.Nonce)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	value, err := utils.BigIntDecode(assetTx.Value)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	callData, err := utils.HexToBytes(assetTx.CallData)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	feeAmount, err := utils.BigIntDecode(assetTx.FeeAmount)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	sig, err := utils.HexToBytes(assetTx.Sig)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	sigKeyIndex, err := utils.BigIntDecode(assetTx.SigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	tx, err := asset.Execute(auth, nonce, assetTx.To, value, callData, assetTx.FeeToken, feeAmount, sig, sigKeyIndex)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
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
