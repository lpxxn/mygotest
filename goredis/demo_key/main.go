package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		//Addr:     "192.168.3.212:6379",
		Password: "", // no password set
		//DB:       3,  // use default DB
	})
	var t1 int64 = 10
	fmt.Println(time.Duration(t1) * time.Second)
	const testKey1 = "testkey1"
	//client.Set(testKey1, "test",10 * time.Second)
	//client.Set(testKey1, "test",10 * time.Minute)

	ttl, err := client.TTL(context.Background(), testKey1).Result()
	// -2s <nil>
	fmt.Println(ttl, err, int64(ttl/time.Second), "s")

	r, err := client.Del(context.Background(), "abcder_dfa").Result()
	fmt.Println("Del :", r, err)

	ctx := context.Background()
	tr, err := client.Set(ctx, "user:1:1:1", "aaaa", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(tr)
	tr, err = client.Set(ctx, "user:1:1:2", "bbbb", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(tr)
	tr, err = client.Set(ctx, "user:1:1:3", "cccc", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(tr)
	sf := func(prefixKey string) {
		fmt.Println("begin scan")
		scanIter := client.Scan(ctx, 0, prefixKey, 0).Iterator()
		for scanIter.Next(ctx) {
			fmt.Println("key: ", scanIter.Val())
		}
		if err := scanIter.Err(); err != nil {
			panic(err)
		}
		fmt.Println("end scan")
	}
	userPrefix := "user:1:*"
	sf(userPrefix)

	// delete
	scanIter := client.Scan(ctx, 0, userPrefix, 0).Iterator()
	for scanIter.Next(ctx) {
		if err := client.Del(ctx, scanIter.Val()).Err(); err != nil {
			fmt.Errorf("Del err: %#v \n", err)
		}
	}
	if err := scanIter.Err(); err != nil {
		panic(err)
	}

	sf(userPrefix)

}
