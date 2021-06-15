package main

import (
	"encoding/json"
	"testing"
)

func TestJson1(t *testing.T) {
	type A struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	type A1 struct {
		A
		ID string `json:"id"`
	}
	a1 := A1{
		A:  A{ID: 123, Name: "li"},
		ID: "abcdef",
	}
	jBody, err := json.Marshal(a1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jBody))

	rev1 := &A1{}
	if err := json.Unmarshal(jBody, rev1); err != nil {
		t.Fatal(err)
	}
	t.Log(rev1)
	a := A{ID: 123, Name: "li"}
	jBody2, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jBody2))
	if err := json.Unmarshal(jBody2, rev1); err != nil {
		t.Fatal(err)
	}
	t.Log(rev1)

}
