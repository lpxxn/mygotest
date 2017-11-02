package main

import (
	"fmt"
	"github.com/mygotest/protocoldemo/demo1/protos"
)

func main() {
	p := tutorial.Person{
		Name: "li",
		Id:   123,
		Phones: []*tutorial.Person_PhoneNumber{
			{Number: "abcdef", Type: tutorial.Person_HOME},
		},
		Email: "lp@1.com",
	}
	fmt.Println(p)
	fmt.Println(p.Phones[0].Type)
	fmt.Println(int(p.Phones[0].Type))
}
