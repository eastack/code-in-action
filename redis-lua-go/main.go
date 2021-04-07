package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	rdb:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	ctx := context.Background()
	rdb.Set(ctx, "hello", "world", time.Duration(10) * time.Minute)

	val, err := rdb.Get(ctx, "hello").Result()

	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val == "":
		fmt.Println("value is empty")
	}


	fmt.Println(val)
}
