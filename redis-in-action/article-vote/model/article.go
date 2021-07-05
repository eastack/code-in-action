package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"article-vote/common"
	"article-vote/util"

	"github.com/go-redis/redis/v8"
)

type Article struct {
	Id       int    `redis:"id"`
	Title    string `redis:"title"`
	AuthorId int    `redis:"authorId"`
	Time     int64  `redis:"time"`
	Votes    int    `redis:"votes"`
	Link     string `redis:"link"`
}

func (a Article) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Article) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, a)
}

type ArticleService interface {
	ArticleVote(ctx context.Context, articleId string, user string)
	PostArticle(ctx context.Context, user string, title string, link string) string
	AddRemoveGroups(ctx context.Context, articleId int, toAdd []string, toRemove []string)
	GetGroupArticles(ctx context.Context, group string, order string, page int64) []map[string]string
}

type ArticleRepo struct {
	Conn *redis.Client
}

func NewArticleRepo(conn *redis.Client) *ArticleRepo {
	return &ArticleRepo{Conn: conn}
}

const scoreZSet = "score:"
const timeZSet = "time:"

func (r *ArticleRepo) PostArticle(
	ctx context.Context,
	userId int,
	title string,
	link string,
) Article {
	// 获取文章ID
	articleId := strconv.Itoa(int(r.Conn.Incr(ctx, "article:").Val()))

	// 发布文章
	now := time.Now().Unix()
	article := Article{
		Id:       util.S2i(articleId),
		Title:    title,
		AuthorId: userId,
		Time:     now,
		Votes:    1,
		Link:     link,
	}
	articleHashKey := "article:" + articleId
	r.Conn.HSet(ctx, articleHashKey, "id", article.Id)
	r.Conn.HSet(ctx, articleHashKey, "title", article.Title)
	r.Conn.HSet(ctx, articleHashKey, "authorId", article.AuthorId)
	r.Conn.HSet(ctx, articleHashKey, "time", article.Time)
	r.Conn.HSet(ctx, articleHashKey, "votes", article.Votes)
	r.Conn.HSet(ctx, articleHashKey, "link", article.Link)
	r.Conn.ZAdd(ctx, scoreZSet, &redis.Z{Score: float64(now + common.VOTE_SCORE), Member: articleHashKey})
	r.Conn.ZAdd(ctx, timeZSet, &redis.Z{Score: float64(now), Member: articleHashKey})

	// 自己默认给自己投一票
	voted := "voted:" + articleId
	r.Conn.SAdd(ctx, voted, userId)
	r.Conn.Expire(ctx, voted, common.ONE_WEEK_IN_SECONDS*time.Second)

	return article
}

func (r *ArticleRepo) ArticleVote(
	ctx context.Context,
	articleId int,
	userId int,
) {
	article := fmt.Sprintf("article:%d", articleId)

	// 一周前的文章无法进行投票了
	cutoff := time.Now().Unix() - common.ONE_WEEK_IN_SECONDS
	if r.Conn.ZScore(ctx, "time:", article).Val() < float64(cutoff) {
		return
	}

	// 尝试投票
	votedSet := fmt.Sprintf("voted:%d", articleId)
	if r.Conn.SAdd(ctx, votedSet, userId).Val() != 0 {
		r.Conn.ZIncrBy(ctx, scoreZSet, common.VOTE_SCORE, article)
		r.Conn.HIncrBy(ctx, article, "votes", 1)
	}
}

func (r *ArticleRepo) GetArticles(
	ctx context.Context,
	page int64,
	order string,
) []map[string]string {
	// 分页参数
	if order == "" {
		order = "score:"
	}
	var perPage int64 = 3
	start := (page - 1) * perPage
	end := start + perPage - 1

	// 按排序规则查询文章ID
	ids := r.Conn.ZRevRange(ctx, order, start, end).Val()

	// 获取详细数据
	var articles []map[string]string
	for _, id := range ids {
		articleData := r.Conn.HGetAll(ctx, id).Val()
		articles = append(articles, articleData)
	}

	return articles
}

func (r *ArticleRepo) AddRemoveGroups(
	ctx context.Context,
	articleId int,
	toAdd []string,
	toRemove []string,
) {
	article := fmt.Sprintf("article:%d", articleId)
	for _, group := range toAdd {
		r.Conn.SAdd(ctx, "group:"+group, article)
	}
	for _, group := range toRemove {
		r.Conn.SRem(ctx, "group:"+group, article)
	}
}

func (r *ArticleRepo) GetGroupArticles(
	ctx context.Context,
	group string,
	order string,
	page int64,
) []map[string]string {
	if order == "" {
		order = "score:"
	}
	key := order + group
	if r.Conn.Exists(ctx, key).Val() == 0 {
		res := r.Conn.ZInterStore(
			ctx, key,
			&redis.ZStore{
				Aggregate: "MAX",
				Keys:      []string{"group:" + group, order},
			},
		).Val()
		if res <= 0 {
			log.Println("ZInterStore return 0")
		}
		r.Conn.Expire(ctx, key, 60*time.Second)
	}
	return r.GetArticles(ctx, page, key)
}
