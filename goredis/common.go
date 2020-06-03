package goredis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

func RClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,  // use default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	return client, err
}
