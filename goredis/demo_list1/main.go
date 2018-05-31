package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	byteValue := []byte("abcdef")

	cmd := client.LPush("a", byteValue)

	i, err := cmd.Result()
	fmt.Print(i, err)

	client.RPush("a", "aaaa")
	client.RPush("b", "aaaa", "ccccc", "ddddd")
}

