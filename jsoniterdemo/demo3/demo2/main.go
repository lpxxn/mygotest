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

type Rsp struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

func main() {
	b := &B{T1: 2, A: A{T1: "abc"}}
	rsp1 := &Rsp{Code:1, Data: b}
	revRsp, err := json.Marshal(rsp1)
	if err != nil {
		panic(err)
	}
	strRev := string(revRsp)
	fmt.Println(strRev)
	rsp2 := &Rsp{}
	err = json.Unmarshal(revRsp, rsp2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rsp2)
	byt, err := json.Marshal(rsp2.Data)
	if err != nil {
		panic(err)
	}
	fmt.Println("byt : ", string(byt))
	revB := &B{}
	UnmarshalB1(byt, revB)

	UnmarshalB2(byt, revB)
}

func UnmarshalB1(rsp []byte, b *B) {
	byt, err := json.Marshal(rsp)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byt, b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b)
}

func UnmarshalB2(rsp []byte, s interface{}) {
	byt, err := json.Marshal(rsp)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byt, s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", s)
}