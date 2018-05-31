package main

import (
	"time"
	"github.com/micro/go-micro"
	"github.com/mygotest/gomicrodemo/demo2/proto"
	"context"
	"fmt"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-plugins/selector/cache"
)

func main() {

	tlsInfo := transport.TLSInfo{
		CertFile:      "/Users/li/certs/s1.pem",
		KeyFile:       "/Users/li/certs/s1-key.pem",
		TrustedCAFile: "/Users/li/certs/etcd-root-ca.pem",
	}

	tls, err := tlsInfo.ClientConfig()

	if err != nil {
		fmt.Println(err)
	}

	registry := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"https://127.0.0.1:2379",
			"https://127.0.0.1:22379",
			"https://127.0.0.1:32379",
		}
		op.TLSConfig = tls
		op.Secure = true
	})

	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"), micro.Registry(registry))
	service.Init()

	service.Client().Init(
		client.Retries(3),
		client.Selector(cache.NewSelector(selector.Registry(registry),)),
	)
	// Create new greeter client
	gr := greeter.NewGreeterService("greeter", service.Client())



	ticker := time.NewTicker(time.Second)
	for{
		// Call the greeter
		rsp, err := gr.Hello(context.TODO(), &greeter.HelloRequest{Name: "lp !!!"})
		if err != nil {
			fmt.Println(err)
		}
		// Print response
		fmt.Println(rsp.Greeting)
		<- ticker.C
	}


}
