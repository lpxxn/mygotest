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
	UnmarshalB1(jsonB, revB)
	UnmarshalB2(jsonB, revB)
}

func UnmarshalB1(rsp []byte, b *B) {
	err := json.Unmarshal(rsp, b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b)
}

func UnmarshalB2(rsp []byte, b interface{}) {
	err := json.Unmarshal(rsp, b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b)
}

