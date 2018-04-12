package main

import (
	"fmt"
	"reflect"
)

type T1 []string
type T2 []string

func main() {

	f0 := []string{	}
	f1 := T1{}
	f2 := T2{}

	fmt.Println(reflect.TypeOf(f0)) // []string
	fmt.Println(reflect.TypeOf(f1))	// main.T1
	fmt.Println(reflect.TypeOf(f2))	// main.T2

	f1 = f0
	f0 = f1
	f0 = f2


	// error
	// f1 = f2


}
