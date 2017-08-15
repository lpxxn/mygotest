package main

import (
	"fmt"
	"reflect"
)

func PrintParam(param string) {
	fmt.Printf("Hello %s \n", param)
}

func PrintParams(param1, param2 string) {
	fmt.Printf("Now Hello %s and %s \n", param1, param2)
}

func Invoke(fn interface{}, args ...string) {
	v := reflect.ValueOf(fn)
	rargs := make([]reflect.Value, len(args))
	for i, a := range  args {
		rargs[i] = reflect.ValueOf(a)
	}
	v.Call(rargs)
}

func main()  {
	Invoke(PrintParam, "Li")
	Invoke(PrintParams, "Peng", "Na")
}
