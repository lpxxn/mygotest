package mylib1

import (
	"fmt"
	"../testlib"
)

type TestData struct{
	Name string `json:"name"`

}

func ReadFun() (string, error) {
	fmt.Println("return string and err")
	testlib.PrintLnThing("test in readutil.go ")
	return "hello ", nil

}
