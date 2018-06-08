package main

import (
	"context"
	"fmt"

	hello "github.com/mygotest/gomicrodemo/greeter/srv/proto/hello"
	"github.com/micro/go-micro"
	"github.com/mygotest/gomicrodemo/greeter/srv/proto/common"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

	// Make request
	rsp, err := cl.Hello(context.Background(), &a_b_common.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
	fmt.Println(rsp.Values[0])
	fmt.Println(rsp.Header)
}
