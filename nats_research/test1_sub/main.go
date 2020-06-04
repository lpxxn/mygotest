package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/nats-io/go-nats"
)

func main() {
	localHost := "nats://192.168.10.208:4222"
	//nc, err := nats.Connect(nats.DefaultURL)
	nc, err := nats.Connect(localHost)
	if err != nil {
		panic(err)
	}

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	fmt.Println("begin subscribe")
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("ending.....")
}

/*
在断网的情况下，可能会丢信息，目前测试，断网，发了3条，等了不到一分钟，再联网，只收到了2条数据
可以直接连 nat-server和 nats-streaming-server

*/
