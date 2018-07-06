package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.212:6379",
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
