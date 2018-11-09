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
	rev, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	strJson1 := string(rev)
	fmt.Println(strJson1)

	b2 := &B{}
	err = json.Unmarshal(rev, b2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b2)

	fmt.Println("----------------")
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
	revB := &B{}
	err = json.Unmarshal(byt, revB)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", revB)


}

