package main

import "fmt"

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}


func A() *int {
	return nil
}

func main() {
	a := A()
	Foo(a)
}