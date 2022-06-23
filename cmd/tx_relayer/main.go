package main

import (
	"fmt"
	"tx_relayer/internel/configs"
	"tx_relayer/internel/logger"
	"tx_relayer/internel/methods"
	"tx_relayer/internel/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

var log = logger.SugarLogger

func main() {
	r := gin.Default()

	log.Info("Asset tx relayer Starting")
	r.Use(middlewares.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	conf, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	handler, err := methods.NewAssetTxHandler(conf)
	if err != nil {
		panic(err)
	}

	r.GET("/chain_id", handler.GetChainId)
	r.GET("/asset_address/:register_email", handler.GetAssetAddress)
	r.POST("/tx/transfer_native", handler.TransferNative)
	r.POST("/tx/transfer_token", handler.TransferToken)
	r.POST("/tx/execute", handler.Execute)

	log.Info("Asset tx relayer Starting")
	fmt.Println("Asset tx relayer Started")

	err = r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
