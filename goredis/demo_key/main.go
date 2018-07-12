package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

func main()  {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.212:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})
	var t1 int64 = 10
	fmt.Println(time.Duration(t1) * time.Second)
	const testKey1 = "testkey1"
	//client.Set(testKey1, "test",10 * time.Second)
	//client.Set(testKey1, "test",10 * time.Minute)

	ttl, err :=client.TTL(testKey1).Result()
	// -2s <nil>
	fmt.Println(ttl, err, int64(ttl/time.Second), "s")

	r, err := client.Del("abcder_dfa").Result()
	fmt.Println("Del :", r, err)
}
