package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	redisGetSet()
	redisScript()
}

func redisGetSet() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	rdb.Set(ctx, "hello", "world", time.Duration(10)*time.Minute)

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

func redisScript() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	IncrByAge := redis.NewScript(`
        if redis.call("GET", KEYS[1]) ~= false then
        	return redis.call("INCRBY", KEYS[1], ARGV[1])
        end
        return false`)

	n, err := IncrByAge.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

	err = rdb.Set(ctx, "xx_counter", "40", 0).Err()
	if err != nil {
		panic(err)
	}

	n, err = IncrByAge.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

	redis.Scripter.ScriptLoad()
}
