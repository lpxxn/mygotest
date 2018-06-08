package main

import (
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
	"context"

	hello "github.com/mygotest/gomicrodemo/greeter/srv/proto/hello"
	"github.com/micro/go-micro"
	"log"
	"time"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	registry := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://192.168.3.34:2379",
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
