package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type A struct {
	Name string `json:"name"`
}
type AList []*A

func TestStructArray(t *testing.T) {
	// "[]"
	a1 := AList{}
	fmt.Printf("%#v \n", a1)
	b1, _ := json.Marshal(a1)
	fmt.Printf("%#v \n", string(b1))

	// "[]"
	a2 := make(AList, 0, 1)
	fmt.Printf("%#v \n", a2)
	b2, _ := json.Marshal(a2)
	fmt.Printf("%#v \n", string(b2))

	// "null"
	var a3 AList
	fmt.Printf("%#v \n", a3)
	b3, _ := json.Marshal(a3)
	fmt.Printf("%#v \n", string(b3))

}
