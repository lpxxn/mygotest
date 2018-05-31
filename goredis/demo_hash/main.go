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

	multiplefields(client)
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


func multiplefields(client *redis.Client) {
	client.HSet("user:1", "name", "li")
	client.HSet("user:1", "age", 10)

	client.HMSet("user:2", map[string]interface{}{"name": "peng", "age": 18})

	u1, err := client.HGetAll("user:1").Result()
	fmt.Println("u1: ", u1, "  err: ", err)

	u2age, err := client.HGet("user:2", "age").Result()

	fmt.Println("u2age: ", u2age, "  err: ", err)


}