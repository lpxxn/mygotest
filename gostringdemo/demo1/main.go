package main

import (
	"encoding/json"
	"fmt"
)

type Atest struct {
	Name string
	Age  int
}

func (a Atest) String() string {
	body, _ := json.Marshal(a)
	return string(body)
}

func main() {
	a := Atest{Name: "li", Age: 18}
	s := fmt.Sprintf("%v", a)
	fmt.Println(s)
}
