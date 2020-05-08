package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/mygotest/goredis"
)

func main() {
	client, err := goredis.RClient()
	if err != nil {
		panic(err)
	}
	rev, _ := client.Ping().Result()
	fmt.Println(rev)
	testBit1 := "testBit1"
	client.Del(testBit1)
	defer client.Del(testBit1)
	_, err = client.SetBit(testBit1, 0, 1).Result()
	if err != nil {
		panic(err)
	}

	bitRev, err := client.GetBit(testBit1, 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(bitRev)
	_, err = client.SetBit(testBit1, 200, 1).Result()
	if err != nil {
		panic(err)
	}
	v, _ := client.Get(testBit1).Result()
	fmt.Println([]byte(v), " len: ", len(v)) // len 26 byte  26 * 8 = 208bit

	_, err = client.SetBit(testBit1, 11250, 1).Result()
	if err != nil {
		panic(err)
	}
	v, _ = client.Get(testBit1).Result()
	fmt.Println([]byte(v), " len: ", len(v)) // 1407 1407*8=11256  11250

	bitRev, err = client.GetBit(testBit1, 11250).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(bitRev)

	bitRev, err = client.GetBit(testBit1, 123231250).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(bitRev)

	notExistsKey := "some:not:exists"
	_, err = client.Del(notExistsKey).Result()
	fmt.Println("delete not exists key error: ", err)
	err = client.Get(notExistsKey).Err()
	fmt.Println(err)
	if err == redis.Nil {
		fmt.Println("redis key not exists")
	}
	bitRev, err = client.GetBit(notExistsKey, 100).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(bitRev)
}
