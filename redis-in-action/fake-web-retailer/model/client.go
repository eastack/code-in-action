package model

import (
	"context"
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fake-web-retailer/common"
	"fake-web-retailer/repository"
	"log"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	Ctx  context.Context
	Conn *redis.Client
}

func NewClient(conn *redis.Client) *Client {
	return &Client{Conn: conn, Ctx: context.Background()}
}

func (r *Client) CheckToken(token string) string {
	return r.Conn.HGet(r.Ctx, "login:", token).Val()
}

func (r *Client) UpdateToken(token, user, item string) {
	timestamp := time.Now().Unix()
	r.Conn.HSet(r.Ctx, "login:", token, user)
	r.Conn.ZAdd(r.Ctx, "recent:", &redis.Z{Score: float64(timestamp), Member: token})
	if item != "" {
		r.Conn.ZAdd(r.Ctx, "viewed:"+token, &redis.Z{Score: float64(timestamp), Member: item})
		r.Conn.ZRemRangeByRank(r.Ctx, "viewed:"+token, 0, -26)
		r.Conn.ZIncrBy(r.Ctx, "viewed:", -1, item)
	}
}

func (r *Client) Reset() {
	r.Conn.FlushDB(r.Ctx)

	common.QUIT = false
	common.LIMIT = 100_000_000
	common.FLAG = 1
}

func (r *Client) CleanSessions() {
	for !common.QUIT {
		size := r.Conn.ZCard(r.Ctx, "recent:").Val()
		if size <= common.LIMIT {
			time.Sleep(1 * time.Second)
			continue
		}

		endIndex := min(size-common.LIMIT, 100)
		tokens := r.Conn.ZRange(r.Ctx, "recent:", 0, endIndex-1).Val()

		var sessionKey []string
		for _, token := range tokens {
			sessionKey = append(sessionKey, token)
		}

		r.Conn.Del(r.Ctx, sessionKey...)
		r.Conn.HDel(r.Ctx, "login:", tokens...)
		r.Conn.ZRem(r.Ctx, "recent:", tokens)
	}
	defer atomic.AddInt32(&common.FLAG, -1)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func (r *Client) AddToCart(session, item string, count int) {
	switch {
	case count <= 0:
		r.Conn.HDel(r.Ctx, "cart:"+session, item)
	default:
		r.Conn.HSet(r.Ctx, "cart:"+session, item, count)
	}
}

func (r *Client) CleanFullSession() {
	for !common.QUIT {
		size := r.Conn.ZCard(r.Ctx, "recent:").Val()
		if size <= common.LIMIT {
			time.Sleep(1 * time.Second)
			continue
		}

		endIndex := min(size-common.LIMIT, 100)
		sessions := r.Conn.ZRange(r.Ctx, "recent:", 0, endIndex-1).Val()

		var sessionKeys []string
		for _, sess := range sessions {
			sessionKeys = append(sessionKeys, "viewed:"+sess)
			sessionKeys = append(sessionKeys, "cart:"+sess)
		}

		r.Conn.Del(r.Ctx, sessionKeys...)
		r.Conn.HDel(r.Ctx, "login:", sessions...)
		r.Conn.ZRem(r.Ctx, "recent:", sessions)
	}
	defer atomic.AddInt32(&common.FLAG, -1)
}

func (r *Client) CacheRequest(request string, callback func(string) string) string {
	if !r.CanCache(request) {
		return callback(request)
	}

	pageKey := "cache:" + hashRequest(request)
	content := r.Conn.Get(r.Ctx, pageKey).Val()

	if content == "" {
		content = callback(request)
		r.Conn.Set(r.Ctx, pageKey, content, 300*time.Second)
	}
	return content
}

func (r *Client) CanCache(request string) bool {
	itemId := extractItemId(request)
	if itemId == "" || isDynamic(request) {
		return false
	}
	rank := r.Conn.ZRank(r.Ctx, "viewed:", itemId).Val()
	return rank != 0 && rank < 10000
}

func (r *Client) ScheduleRowCache(rowId string, delay int64) {
	r.Conn.ZAdd(r.Ctx, "delay:", &redis.Z{Member: rowId, Score: float64(delay)})
	r.Conn.ZAdd(r.Ctx, "schedule:", &redis.Z{Member: rowId, Score: float64(time.Now().Unix())})
}

func (r *Client) CacheRows(rowId string, delay int64) {
	for !common.QUIT {
		next := r.Conn.ZRangeWithScores(r.Ctx, "schedule:", 0, 0).Val()
		now := time.Now().Unix()

		if len(next) == 0 || next[0].Score > float64(now) {
			time.Sleep(50 * time.Microsecond)
			continue
		}

		rowId := next[0].Member.(string)
		delay := r.Conn.ZScore(r.Ctx, "delay:", rowId).Val()
		if delay <= 0 {
			r.Conn.ZRem(r.Ctx, "delay:", rowId)
			r.Conn.ZRem(r.Ctx, "schedule:", rowId)
			r.Conn.Del(r.Ctx, "inv:"+rowId)
		}

		row := repository.Get(rowId)
		r.Conn.ZAdd(r.Ctx, "schedule:", &redis.Z{Member: rowId, Score: float64(now) + delay})
		jsonRow, err := json.Marshal(row)
		if err != nil {
			log.Fatalf("marshal json failed, data is: %v, err is: %v\n", row, err)
		}

		r.Conn.Set(r.Ctx, "inv:"+rowId, jsonRow, 0)
	}
	defer atomic.AddInt32(&common.FLAG, -1)
}
func (r *Client) RescaleViewed(rowId string, delay int64) {
	for !common.QUIT {
		r.Conn.ZRemRangeByRank(r.Ctx, "viewed:", 20_000, -1)
		r.Conn.ZInterStore(r.Ctx, "viewed:", &redis.ZStore{Weights: []float64{0.5}, Keys: []string{"viewed:"}})
		time.Sleep(300 * time.Second)
	}
}
func extractItemId(request string) string {
	parsed, _ := url.Parse(request)
	queryValue, _ := url.ParseQuery(parsed.RawQuery)
	query := queryValue.Get("item")
	return query
}

func isDynamic(request string) bool {
	parsed, _ := url.Parse(request)
	queryValue, _ := url.ParseQuery(parsed.RawQuery)
	for _, v := range queryValue {
		for _, j := range v {
			if strings.Contains(j, "_") {
				return false
			}
		}
	}
	return true
}

func hashRequest(request string) string {
	hash := crypto.MD5.New()
	hash.Write([]byte(request))
	res := hash.Sum(nil)
	return hex.EncodeToString(res)
}
