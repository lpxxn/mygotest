package main

import "fmt"

type Test struct {
	A string
}

func main() {
	t := Test{A :"aaa"}
	addr := fmt.Sprintf("%p", &t)
	fmt.Println(addr)
}
