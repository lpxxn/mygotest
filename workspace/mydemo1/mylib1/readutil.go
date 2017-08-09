package mylib1

import "fmt"

type TestData struct{
	Name string `json:"name"`

}

func ReadFun() (string, error) {
	fmt.Println("return string and err")
	return "hello ", nil

}
