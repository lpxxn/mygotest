package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string
}

func main() {
	t := Test{Name:"li"}
	fmt.Println(reflect.ValueOf(t).Kind() == reflect.Struct)
}
