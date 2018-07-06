package main

import (
	"github.com/go-redis/redis"
	"time"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.212:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})
	const uidSayCount = "saycount:uid"
	v, err := client.Incr(uidSayCount).Result()
	if err == redis.Nil {
		pipe := client.TxPipeline()

		//inc := pipe.Incr("uidSayCount")
		pipe.Expire(uidSayCount, time.Hour)

		// Execute
		//
		//     MULTI
		//     INCR pipeline_counter
		//     EXPIRE pipeline_counts 3600
		//     EXEC
		//
		// using one client-server roundtrip.

		_, err := pipe.Exec()

		if err != nil {
			panic(err)
		}

	} else if err != nil {
		panic(err)
	}

	fmt.Println(v)




}
