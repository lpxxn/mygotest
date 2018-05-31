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
type TestHello struct{}

func (g *TestHello) DoAction(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.RespDesc = "do action: 2 " + req.Name
	return nil
}

func (g *TestHello)SaySomeThing(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.RespDesc = "say something to 2: " + req.Name
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

	/*
			"https://192.168.3.34:2379",
		"https://192.168.3.18:2379",
		"https://192.168.3.110:2379",

 */

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
		micro.Name("tstmicroservice"),
		// Set service registry
		micro.Registry(registry),

		// TTL 发现服务信息存多长时间，过期后被删除
		micro.RegisterTTL(time.Second * 30),
		// 时间间隔是服务应该重新注册的时候，以保留在服务发现中的注册信息
		micro.RegisterInterval(time.Second * 15),
	)

	service.Init()
	micro.NewFunction()
	tsv := new(TestHello)
	proto.RegisterPersonHandler(service.Server(), tsv)
	proto.RegisterSayHandler(service.Server(), tsv)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
