package main

import (
	"fmt"
	"reflect"
)

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func Foo2(x interface{}) {
	if reflect.ValueOf(x).IsNil() {
		fmt.Println("x is empty interface")
		return
	}
	fmt.Println("x is non-empty interface")
}


func A() *int {
	return nil
}

func B() *int {
	var i *int
	return i
}

type T struct {

}

func M() *T {
	var t *T
	return t
}

func main() {
	a := A()
	if a == nil {
		fmt.Println("a is nil")
	}

	Foo(a)
	Foo2(a)

	b := B()
	Foo(b)
	Foo2(b)

	t1:= M()
	if t1 == nil {
		fmt.Println("t1 is nil")
	}

}