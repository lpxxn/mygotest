package main


import (
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro"
	"github.com/coreos/etcd/pkg/transport"
	"fmt"
	"github.com/micro/go-micro/registry"
	"context"
	"github.com/mygotest/gomicrodemo/demo2/proto"
	"time"
)

// we can open more service by modify the Hello Response
type TestGreeter struct{}

func (g *TestGreeter) Hello(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	rsp.Greeting = "Hello 3" + req.Name
	return nil
}


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


	service := micro.NewService(
		// Set service name
		micro.Name("greeter"),
		// Set service registry
		micro.Registry(registry),

		// TTL 发现服务信息存多长时间，过期后被删除
		micro.RegisterTTL(time.Second * 30),
		// 时间间隔是服务应该重新注册的时候，以保留在服务发现中的注册信息
		micro.RegisterInterval(time.Second * 15),
	)

	service.Init()
	micro.NewFunction()

	greeter.RegisterGreeterHandler(service.Server(), new(TestGreeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
