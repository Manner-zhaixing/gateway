package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化路由
	r := gin.Default()
	api := r.Group("/api")
	{
		controller := handlers.NewRpcServiceController(rpcServiceService)
		api.POST("/rpcservice", controller.CreateRpcService)
		api.GET("/rpcservice/:id", controller.GetRpcService)
		api.PUT("/rpcservice/:id", controller.UpdateRpcService)
		api.DELETE("/rpcservice/:id", controller.DeleteRpcService)
	}

	r.Run(":8080")
}
