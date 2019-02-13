package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string
}

func (t Test) GetName() string {
	return t.Name
}

func (t Test) Cal(a, b int) int {
	fmt.Printf("a: %d  b: %d \n", a, b)
	return a + b
}

type Tester interface {
	GetName() string
}

func main() {
	t := Test{Name: "li"}
	tv := reflect.ValueOf(t)
	fmt.Println(tv.Kind() == reflect.Struct)

	interType := reflect.TypeOf((*Tester)(nil)).Elem()
	if tv.Type().Implements(interType) {
		n := tv.Interface().(Tester).GetName()
		fmt.Println("get t name: ", n)
	}

	fmt.Println("------")
	//
	typ := reflect.TypeOf(t)
	//hdlr := reflect.ValueOf(t)
	//name := reflect.Indirect(hdlr).Type().Name()

	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mt := method.Type
		fmt.Println("method name", method.Name)
		numIn := mt.NumIn()
		fmt.Println("numIn: ", numIn)
		for idx := 0; idx < numIn; idx++ {
			argT := mt.In(idx)
			fmt.Println("name: ", argT.Name())
		}

		if method.Name == "Cal" {
			rev := method.Func.Call([]reflect.Value{reflect.ValueOf(t), reflect.ValueOf(1), reflect.ValueOf(2)})
			fmt.Println("rev :", rev[0].Interface())
		}
	}
}
