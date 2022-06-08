package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,  // use default DB
	})
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	byteValue := []byte("abcdef")

	cmd := client.LPush(ctx, "a", byteValue)

	i, err := cmd.Result()
	fmt.Println(i, err)

	//client.RPush("a", "aaaa")
	//client.RPush("b", "aaaa", "ccccc", "ddddd")

	r, err := client.LPop(ctx, "b").Result()
	if err == redis.Nil {
		fmt.Println("not find")
	}

	fmt.Println(err, r)

	l, err := client.LLen(ctx, "a").Result()
	fmt.Println(l)
	type a1 struct {
		ID string `json:"ID"`
	}
	type a2 struct {
		ID int64 `json:"ID"`
		a1
	}
	v := a2{
		ID: 123,
		a1: a1{ID: "abcdef"},
	}
	b, _ := json.Marshal(v)
	fmt.Println(string(b))
}

var Client *redis.Client
var script string = `
		local value = redis.call("Get", KEYS[1])
		print("当前值为 " .. value);
		if( value - KEYS[2] >= 0 ) then
			local leftStock = redis.call("DecrBy" , KEYS[1],KEYS[2])
   			print("剩余值为" .. leftStock );
			return leftStock
		else
			print("数量不够，无法扣减");
			return value - KEYS[2]
		end
		return -1
	`
var luaHash string
