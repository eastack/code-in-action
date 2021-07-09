package main

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"fake-web-retailer/common"
	"fake-web-retailer/model"

	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

func TestLoginCookies(t *testing.T) {
	conn := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	client := model.NewClient(conn)
	token := uuid.NewV4().String()
	username := "eastack"

	t.Run("Test UpdateToken", func(t *testing.T) {
		client.UpdateToken(token, username, "itemX")
		t.Log("We just logged-in/update token: \n", token)
		t.Log("For user: ", username, "\n")
		t.Log("\nWhat user do wo get when we look-up that token?")
		r := client.CheckToken(token)
		t.Log("username:", r)
		t.Helper()
		if r != username {
			t.Errorf("want get %v, actual get %v\n", r, username)
		}

		t.Log("Let's drop the maximum number of cookies to 0 to clean them out\n")
		t.Log("We will start a thread to do the cleaning, while we stop it later\n")

		common.LIMIT = 0
		go client.CleanSessions()
		time.Sleep(1 * time.Second)
		common.QUIT = true
		time.Sleep(2 * time.Second)

		client.Reset()
	})
}

func TestUrl(t *testing.T) {
	baiduUrl := "https://baidu.com/?item=abc&keyword=test"
	parsed, _ := url.Parse(baiduUrl)
	queryValue, _ := url.ParseQuery(parsed.RawQuery)
	for k, v := range queryValue {
		println(k)
		for ki, vi := range v {
			println(ki)
			println(vi)
		}
	}
}

func TestZRank(t *testing.T) {
	t.Run("Test zset rank", func(t *testing.T) {
		conn := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})

		v := conn.ZRank(context.Background(), "z", "sony").Val()
		fmt.Println(v)
		conn.Close()
	})
}
