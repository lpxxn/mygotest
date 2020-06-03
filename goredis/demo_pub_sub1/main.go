package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.10.131:6379",
		Password: "", // no password set
	})
	ctx, cancel := context.WithCancel(context.Background())
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	keys, err := client.Keys(ctx, "*").Result()
	fmt.Println(keys, err)

	ps := client.Subscribe(ctx, "top1")
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		fmt.Println("cancel ctx")
		cancel()
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ending ....")
			return
		default:
		}
		ctxT, cancelT := context.WithTimeout(ctx, time.Second*5)
		msg, err := ps.ReceiveMessage(ctxT)
		if err != nil {
			if os.IsTimeout(err) {
				continue
			}
			fmt.Printf("ReceiveMessage err := %#v", err)
			continue
		}
		cancelT()
		fmt.Println("msg: ", msg.String())
	}
}

/*
当网络终断后 再次连上后，会取得断开前的数据
SUBSCRIBE someotherchannel

PUBLISH top1 abc
PUBLISH someotherchannel ipsum


1.在RedisClient 内部维护了一个pubsub_channels的Channel列表，记录了此客户端所订阅的频道

2.在Server服务端，同样维护着一个类似的变量叫做，pubsub_channels，这是一个dict字典变量，每一个Channel对应着一批订阅了此频道的Client，也就是Channel-->list of Clients

3.当一个Client publish一个message的时候，会先去服务端的pubsub_channels找相应的Channel，遍历里面的Client，然后发送通知，即完成了整个发布订阅模式。
*/
