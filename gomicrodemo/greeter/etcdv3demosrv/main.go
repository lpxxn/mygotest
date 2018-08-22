package main

import (
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
	"context"

	hello "github.com/mygotest/gomicrodemo/greeter/srv/proto/hello"
	"github.com/mygotest/gomicrodemo/greeter/srv/proto/common"
	"github.com/micro/go-micro"
	"log"
	"time"
	"fmt"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *a_b_common.Request, rsp *a_b_common.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	s1 := string("start...")
	fmt.Println(s1)
	registry := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
		}
	})

	service := micro.NewService(
		micro.Name("test34"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),


		micro.Registry(registry),
	)


	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
