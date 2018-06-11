package main

import (
	"github.com/json-iterator/go"
	"fmt"
)

func main() {
	p := fmt.Println
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	p(jsoniter.Get(val, "Colors", 0).ToString())
	p(jsoniter.Get(val, "Name").ToString())
	if jsoniter.Get(val, "Type").ToString() == "" {
		p("empty")
	}

}
