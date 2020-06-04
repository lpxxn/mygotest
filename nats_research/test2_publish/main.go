package main

import (
	"fmt"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	//localHost := "nats://192.168.10.208:4222"
	nc, err := nats.Connect(nats.DefaultURL)
	//nc, err := nats.Connect(localHost)
	if err != nil {
		panic(err)
	}
	if err := nc.Publish("foo", []byte(fmt.Sprintf("Hello World %d", time.Now().Unix()))); err != nil {
		panic(err)
	}
	fmt.Println("end publish")
}

/*
可以直接连 nat-server和 nats-streaming-server

*/
