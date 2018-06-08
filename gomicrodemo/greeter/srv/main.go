package main

import (
	"log"
	"time"

	hello "github.com/mygotest/gomicrodemo/greeter/srv/proto/hello"
	"github.com/micro/go-micro"

	"context"
	"github.com/mygotest/gomicrodemo/greeter/srv/proto/common"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *a_b_common.Request, rsp *a_b_common.Response) error {
	log.Print("Received Say.Hello request : ", req)
	rsp.Msg = "Hello " + req.Name
	rsp.Header =  map[string]*a_b_common.Pair{"abc": &a_b_common.Pair{Key: "abc", Values: []string {"1", "2"}}}
	rsp.Values = []string{"a"}
	rsp.Type = 2
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
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

/*
	protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. ./*.proto


或者
	micro api  --handler=rpc

	curl -d 'service=go.micro.srv.greeter' \
    -d 'method=Say.Hello' \
    -d 'request={"name": "Asim Aslam"}' \
    http://localhost:8080/rpc


	curl -d 'service=go.micro.srv.greeter' \
    -d 'method=Say.Hello' \
    -d 'request={"name": "Asim Aslam", "values": [2, 3, 4], "header": {"a": {"key": "k", "values": ["aaaa"]}}, "type": 1}' \
    http://localhost:8080/rpc



	curl -d 'service=li.peng.a' \
    -d 'method=Foo.Bar' \
    -d 'request={}' \
    http://localhost:8080/rpc


//
consul agent -dev


/Users/li/go/src/github.com/mygotest/gomicrodemo/greeter/srv/proto

protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. ./hello/*.proto

protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. ./common/*.proto
protoc --proto_path=$GOPATH/src:. --go_out=. ./common/*.proto

因为common里没有micro的接口，可以不用输出 micro_out

--proto_path 是根目录，import的文件将从要目录下开如寻找

 */