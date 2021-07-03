package model

import (
	"context"
	"strconv"
	"time"

	"article-vote/common"

	"github.com/go-redis/redis/v8"
)

type Article interface {
	ArticleVote(ctx context.Context, articleId string, user string)
	PostArticle(ctx context.Context, user string, title string, link string) string
}

type ArticleRepo struct {
	Conn *redis.Client
}

func NewArticleRepo(conn *redis.Client) *ArticleRepo {
	return &ArticleRepo{Conn: conn}
}

func (r *ArticleRepo) ArticleVote(ctx context.Context, articleId string, user string) {
	articleKey := "article:" + articleId
	cutoff := time.Now().Unix() - common.ONE_WEEK_IN_SECONDS
	if r.Conn.ZScore(ctx, "time:", articleKey).Val() < float64(cutoff) {
		return
	}

	if r.Conn.SAdd(ctx, "voted:"+articleId, user).Val() != 0 {
		r.Conn.ZIncrBy(ctx, "score:", common.VOTE_SCORE, articleKey)
		r.Conn.HIncrBy(ctx, articleKey, "votes", 1)
	}
}

func (r *ArticleRepo) PostArticle(
	ctx context.Context,
	user string,
	title string,
	link string,
) string {
	articleId := strconv.Itoa(int(r.Conn.Incr(ctx, "article:").Val()))

	voted := "voted:" + articleId
	r.Conn.SAdd(ctx, voted, user)
	r.Conn.Expire(ctx, voted, common.ONE_WEEK_IN_SECONDS*time.Second)

	now := time.Now().Unix()
	article := "article:" + articleId
	r.Conn.HSet(ctx, article, map[string]interface{}{
		"title":  title,
		"link":   link,
		"poster": user,
		"time":   now,
		"votes":  1,
	})

	r.Conn.ZAdd(ctx, "score:",
		&redis.Z{Score: float64(now + common.VOTE_SCORE), Member: article})
	r.Conn.ZAdd(ctx, "time:",
		&redis.Z{Score: float64(now), Member: article})

	return articleId
}

//func (r *ArticleRepo) GetArticles(page int64, order string) []map[string]string {
//
//}
