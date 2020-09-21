package main

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
var ctx = context.Background()
func main() {
	y, w := time.Now().ISOWeek()
	fmt.Println(time.Now().Weekday(), "  ", y, w)
	fmt.Printf("%.2f\n", 100.289999)
	fmt.Println(big.NewFloat(100.289999).Text('f', 2))
	f1 := 100.28999
	v := int64(f1 * 100)
	fmt.Println(float64(v) / 100)
	client, _ := RClient()

	ttl, err := RoomDisableUidExpire(client, "111", "333")
	fmt.Println("aa", ttl, err)

	u, err := RoomUidExist(client, "111", "333_lipeng")
	fmt.Println("RoomUidExist", u, err)
	//testSetAdd(client)
	testSetRemove(client, []string{"UqLKGhSm", "eiYyy"})

	l, err := client.SCard(ctx, "t").Result()
	fmt.Println("len : ", l, "  err: ", err)

	a, err := RoomAdminExit(client, "1111", "aa")
	if err == redis.Nil {
		//"a is 0"
		fmt.Println("nil")
	}
	fmt.Println(a, err)
}

func RClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Addr:     "192.168.3.212:6379",
		Password: "", // no password set
		//DB:       2,  // use default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	return client, err
}

func testSetAdd(client *redis.Client) {
	for i := 0; i < 3000; i++ {
		client.SAdd(ctx, "t", randSeq(rand.Intn(10)), randSeq(rand.Intn(15)))
	}
}

func testSetPop(client *redis.Client) {
	client.SPop(ctx, "t")
}

func testSetRemove(client *redis.Client, m []string) {
	cmd := client.SRem(ctx, "t", m)
	i, err := cmd.Result()
	fmt.Println("count: ", i, " err :", err)
	if err != nil {
		panic(err)
	}
}

func RoomAdminExit(client *redis.Client, roomId string, member string) (float64, error) {
	roomName := "test" + roomId

	return client.ZScore(ctx, roomName, member).Result()
}

func RoomUidExist(client *redis.Client, roomId string, members ...interface{}) (int64, error) {
	roomName := "roominfo:" + roomId

	return client.ZRem(ctx, roomName, members...).Result()
}

func RoomDisableUidExpire(client *redis.Client, roomId, uid string) (time.Duration, error) {
	name := "roomdisuid:" + roomId + ":" + uid
	return client.TTL(ctx, name).Result()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
