package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rwrr/oauth2-in-action/controller/client"
	"log"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/client/*")

	router.GET("/", client.Index)
	router.GET("/authorize", client.Authorize)
	router.GET("/callback", client.Callback)
	router.GET("/fetch_resource", client.FetchResource)

	err := router.Run("localhost:9000")
	if err != nil {
		log.Fatalf("Server starting failed: %s", err.Error())
	}
}
