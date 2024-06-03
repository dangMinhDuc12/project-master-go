package redis

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := rdb.Ping(context.Background()).Result()

	 if err != nil {
        panic("Failed to connect to Redis: " + err.Error())
    }
    fmt.Println("Connected to Redis")

		return rdb
}