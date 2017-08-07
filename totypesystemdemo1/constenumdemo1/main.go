package main

import "fmt"

const (
	a = 1 << iota
	b
	c
)

const (
	d = 1 << iota
	e
	_	// skips iota = 2
	g
	h
)

type MyIntEnum int
const (
	i MyIntEnum = iota
	j
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(e, " g: ", g)

	fmt.Println(i, j)
	
	var testIenu MyIntEnum = MyIntEnum(1)

	switch testIenu {
	case i:
		fmt.Println("MyIntEnum i", i)
	case j:
		fmt.Println("MyIntEnum j", j)
	default:
		fmt.Println("default value", testIenu)
	}
}
