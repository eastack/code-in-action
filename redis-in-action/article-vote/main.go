package main

import (
	"time"

	"article-vote/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Article struct {
	title  string
	link   string
	poster int
	time   time.Time
	votes  int
}

func setupRouter() *gin.Engine {
	var rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	articleRepo := model.NewArticleRepo(rdb)

	r := gin.Default()

	r.POST("/article/:id", func(c *gin.Context) {
		user_id := c.Param("id")
		title := c.Query("title")
		link := c.Query("link")
		articleId := articleRepo.PostArticle(c, user_id, title, link)
		c.JSON(200, gin.H{
			"id": articleId,
		})
	})
	r.POST("/article/:id/vote/up", func(c *gin.Context) {
		articleId := c.Param("id")
		userId := c.Query("userId")
		articleRepo.ArticleVote(c, articleId, userId)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
