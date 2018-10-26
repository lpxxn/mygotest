package main

import (
	"fmt"
	"strconv"
)

type TS1 struct {
	A string
	S1 []Ts1Detail
}

type Ts1Detail struct {
	D1 string
}

func (t *TS1) Updata(i int) {
	t.A = "aa" + strconv.Itoa(i)
	t.S1 = []Ts1Detail{{D1: strconv.Itoa(i) + "aaa"}}
}

type Ts2 struct {
	B string
	S2 []TS1
}

func main() {
	t2 := Ts2{B: "B"}


	for i :=0; i < 2; i++ {
		t1 := TS1{A: "li", S1:[]Ts1Detail{{D1: "123"}}}
		t1.Updata(i)
		t2.S2 = append(t2.S2, t1)
		//fmt.Println(t1)
	}
	fmt.Println(t2)

}
