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
}
