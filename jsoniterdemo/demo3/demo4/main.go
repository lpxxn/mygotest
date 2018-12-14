package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	T1 string `json:"t1"`
}

type B struct {
	A
	T1 int `json:"t1"`
}


func main() {
	b := &B{T1: 2, A: A{T1: "abc"}}
	jsonB, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	revB := &B{}
	// have error! why
	err = json.Unmarshal(jsonB, revB)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", revB)

}


