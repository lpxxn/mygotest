package main

import (
	"fmt"
	"time"
)

func main() {
	var v int64 = 10
	t := time.Duration(v)
	fmt.Println(t)
	s := Student{}
	s.Say()
}


type People struct {
	Name string
	Age int32
}

func (p People) SayHello() {
	fmt.Println("Hello World")
}

func (p People) Say()  {
	p.SayHello()
}

type Student struct {
	People
}

func (s Student) SayHello() {
	fmt.Println("Hello World Im Student")
}