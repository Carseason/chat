package main

import (
	"chat/handle"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	v1 := router.Group("/v1") //组路由
	{
		v1.GET("/websocket", handle.HandleWebSocket)
	}
	go handle.HandleMessages()
	router.Run(":8082")
}
