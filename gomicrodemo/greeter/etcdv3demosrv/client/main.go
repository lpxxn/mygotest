package main

import (
	"context"

	hello "github.com/mygotest/gomicrodemo/greeter/srv/proto/hello"
	"github.com/mygotest/gomicrodemo/greeter/srv/proto/common"
	"github.com/micro/go-micro"
	"fmt"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
	"time"
	"os"
	"os/signal"
)

func main(){

	registry := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://192.168.3.34:2379",
		}
	})
	// create a new service
	service := micro.NewService(
		micro.Registry(registry),
	)

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := hello.NewSayService("test34", service.Client())


	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, os.Interrupt)


	tick := time.Tick(time.Second * 3)

	for {
		select {
		case <-tick:
			// Make request
			rsp, err := cl.Hello(context.Background(), &a_b_common.Request{
				Name: "John",
			})
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(rsp)
			fmt.Println(rsp.Values)
			fmt.Println(rsp.Header)
		case <- stopCh:
			return
		}
	}


}
