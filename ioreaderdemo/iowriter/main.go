package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func (p *Person) Write(w io.Writer) (n int, err error) {
	b, _ := json.Marshal(p)
	return w.Write(b)
}

func main() {
	me := Person{
		Id:   1,
		Name: "H",
		Age:  10,
	}
	var b bytes.Buffer
	me.Write(&b)
	fmt.Println(b.String())
}
