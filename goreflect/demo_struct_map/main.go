package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Weight float64
}

func main() {
	p := new(Person)
	p2 := &Person{"allan", 18, 62.3}

	TryIsValid(p)
	TryIsValid(p2)

	TryIsZero(p)
	TryIsZero(p2)
}

func TryIsValid(p interface{}) {
	v := reflect.ValueOf(p)
	ind := reflect.Indirect(v)
	num := ind.NumField()
	fmt.Println("****************************")
	for i := 0; i < num; i++ {
		fv := ind.Field(i)
		if fv.IsValid() {
			fmt.Println("Valid:", fv)
		} else {
			fmt.Println("Not a valid field value: ", fv)
		}
	}
	fmt.Println("****************************\n")
}

func TryIsZero(p interface{}) {
	v := reflect.ValueOf(p)
	ind := reflect.Indirect(v)
	num := ind.NumField()
	fmt.Println("****************************")
	for i := 0; i < num; i++ {
		fv := ind.Field(i)
		if IsZero(fv) {
			fmt.Println("Zero:", fv)
		} else {
			fmt.Println("Not Zero: ", fv)
		}
	}
	fmt.Println("****************************\n")
}

func IsZero(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(reflect.TypeOf(v.Interface())).Interface())
}
