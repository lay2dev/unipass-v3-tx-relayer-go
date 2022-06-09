package main

import (
	"asset_forwarder/internel/logger"
	"asset_forwarder/internel/methods"
	"asset_forwarder/internel/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

var log = logger.SugarLogger

func main() {
	r := gin.Default()

	log.Info("Asset Forwarder Starting")
	r.Use(middlewares.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handler, err := methods.NewAssetTxHandler()
	if err != nil {
		panic(err)
	}

	r.GET("/chain_id", handler.GetChainId)
	r.GET("/asset_address/:register_email", handler.GetAssetAddress)
	r.POST("/tx/transfer_native", handler.TransferNative)
	r.POST("/tx/transfer_token", handler.TransferToken)
	r.POST("/tx/execute", handler.Execute)

	fmt.Println("Asset Transaction Forwarder Started")
	r.Run() // listen and serve on 0.0.0.0:8080
}
