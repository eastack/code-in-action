package router

import (
	"article-vote/model"
	"article-vote/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func SetupArticleRouter(engine *gin.Engine) *gin.Engine {
	var rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	articleRepo := model.NewArticleRepo(rdb)

	engine.POST("/articles", func(c *gin.Context) {
		userId := util.S2i(c.GetHeader("Authorization"))
		title := c.Query("title")
		link := c.Query("link")
		article := articleRepo.PostArticle(c, userId, title, link)
		c.JSON(200, article)
	})

	engine.GET("/articles", func(c *gin.Context) {
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		order := c.Query("order")
		articles := articleRepo.GetArticles(c, page, order)
		c.JSON(200, articles)
	})

	engine.POST("/articles/:id/vote/up", func(c *gin.Context) {
		articleId := util.S2i(c.Param("id"))
		userId := util.S2i(c.GetHeader("Authorization"))
		articleRepo.ArticleVote(c, articleId, userId)
	})

	return engine
}
