package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/mygotest/tcp_udp/udp_demo/demo1"
)

func main() {
	hostPort := "127.0.0.1:9876"
	client, err := demo1.NewTUPPClientTransport(hostPort)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case <-time.After(time.Second/10):
				if _, err := client.Write([]byte("hello world")); err != nil {
					fmt.Println(err)
				}
				if err := client.Flush(); err != nil {
					var opErr *net.OpError
					if ok := errors.As(err, &opErr); ok {
						fmt.Println(opErr, " is timeout: ", opErr.Timeout(), "  tempary: ", opErr.Temporary())
					} else if errors.Is(err, net.ErrWriteToConnected) {
						fmt.Println("ErrWriteToConnected")
					} else {
						fmt.Println(err)
					}
				}
			}
		}
	}()

	osCh := make(chan os.Signal)
	signal.Notify(osCh, os.Interrupt)
	<-osCh
}
