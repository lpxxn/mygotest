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
	revB2 := &B{}
	// panic
	UnmarshalB1(jsonB, revB2)

}

func UnmarshalB1(rsp []byte, b *B) {
	byt, err := json.Marshal(rsp)
	if err != nil {
		panic(err)
	}
	// have error! why
	err = json.Unmarshal(byt, b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b)
}

