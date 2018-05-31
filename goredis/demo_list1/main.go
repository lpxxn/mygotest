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
	fmt.Println(i, err)

	//client.RPush("a", "aaaa")
	//client.RPush("b", "aaaa", "ccccc", "ddddd")


	r, err := client.LPop("b").Result()
	if err == redis.Nil {
		fmt.Println("not find")
	}

	fmt.Println(err, r)

	l, err := client.LLen("a").Result()
	fmt.Println(l)


}

