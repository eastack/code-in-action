package main

import (
	"article-vote/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.SetupArticleRouter(r)

	r.Run(":8080")
}
