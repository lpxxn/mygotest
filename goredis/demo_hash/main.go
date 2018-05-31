package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	client, err := RClient()
	if nil != err {
		panic(err)
	}
	client.HSet("ma", "f1", "abc")
	client.HMSet("mb", map[string]interface{} {
		"a": "1",
		"b": "2",
	})

	s1, err := client.HGet("ma", "f1").Result()
	fmt.Println(s1, err)

	i, err := client.HDel("ma", "f1").Result()
	fmt.Println(i, err)

	len, err := client.HLen("mb").Result()
	fmt.Println("len: ", len)
}



func RClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client, err
}
