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

type ByteSize float64

const (
	B           = iota + 1                  // ignore first value by assigning to blank identif
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                                   // 1 << (10*2)
	GB                                   // 1 << (10*3)
	TB                                   // 1 << (10*4)
	PB                                   // 1 << (10*5)
	EB                                   // 1 << (10*6)
	ZB                                   // 1 << (10*7)
	YB                                   // 1 << (10*8)
)

type MyIntEnum int
const (
	i MyIntEnum = iota
	j
)

func main() {

	fmt.Println(B, KB, MB, GB)
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
