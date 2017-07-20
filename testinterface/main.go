package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return strings.Join([]string{"Name: ", p.Name, " the Age : ", strconv.Itoa(p.Age)}, "")
}

func main() {
	fmt.Println("hello")
	var iemp interface{}
	desc(iemp)
	iemp = 1
	desc(iemp)
	iemp = "hello"
	desc(iemp)

	v, ok := iemp.(int)
	if ok == true {
		desc(v)
	} else {
		fmt.Println("false")
	}

	v1, ok1 := iemp.(string)
	if ok1 == true {
		desc(v1)
	} else {
	}
	switchType(21)
	switchType("hell")
	switchType(Person{"li", 10})
	switchType(&Person{"peng", 10})

	var p1 *Person = &Person{"li", 10}
	p2 := &Person{"peng", 1}
	fmt.Println(p1, p2)
}

func desc(i interface{}) {
	fmt.Printf("value: %v , type: %T \n", i, i)
}

func switchType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("string %v \n", i)
	case int:
		fmt.Printf("%v , %T \n", i, v)
	default:
		fmt.Printf("%v, %T \n", i, v)
	}
}
