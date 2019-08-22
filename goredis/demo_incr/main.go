package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})
	const uidSayCount = "saycount:uid"
	v, err := client.Incr(uidSayCount).Result()
	fmt.Println("err", err)
	if err != nil {
		panic(err)
	}
	if v == 1 {
		client.Expire(uidSayCount, time.Hour)
	}
	if v > 10 {
		//
	}

	fmt.Println(v)
}
