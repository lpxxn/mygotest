package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/miekg/dns"
	"github.com/mygotest/tcp_udp/udp_demo/mdns_demo1"
)

func main() {
	client, err := mdns_demo1.NewMClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	msgCh := make(chan *dns.Msg, 32)
	go client.Recv(client.IPv4UnicastConn, msgCh)
	go client.Recv(client.IPv4MulticastConn, msgCh)

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	for {
		select {
		case <-exit:
			return
		case <-client.ClosedCh:
			return
		case m := <-msgCh:
			fmt.Println(m.String())
		}
	}
}
