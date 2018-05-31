package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main(){
	client, _ := RClient()
	//testSetAdd(client)
	testSetRemove(client, []string{"UqLKGhSm", "eiYyy"})

	l, err := client.SCard("t").Result()
	fmt.Println("len : ", l, "  err: ", err)
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


func testSetAdd(client *redis.Client) {
	for i := 0; i < 3000; i++ {
		client.SAdd("t", randSeq(rand.Intn(10)), randSeq(rand.Intn(15)))
	}
}

func testSetPop(client *redis.Client) {
	client.SPop("t")
}

func testSetRemove(client *redis.Client, m []string) {
	cmd := client.SRem("t", m)
	i, err := cmd.Result()
	fmt.Println("count: ", i, " err :", err)
	if err != nil {
		panic(err)
	}
}



var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}