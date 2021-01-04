package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rwrr/oauth2-in-action/controller/server"
	"log"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/server/*")

	router.GET("/", server.Index)
	router.GET("/authorize", server.Authorize)
	router.POST("/approve", server.Approve)
	router.POST("/token", server.Token)

	err := router.Run("localhost:9001")
	if err != nil {
		log.Fatalf("Server starting failed: %s", err.Error())
	}
}
