package main

import (
	"fmt"
)

type Ivs1 []int

func (iv *Ivs1) Append(d ...int) {
	*iv = append([]int(*iv), d...)
}

func AppendIvs(s Ivs1, d ...int) {
	s = append(s, d...)
}

func changeIvs(s Ivs1) {
	s[0] = 111111
}
func main() {
	var i1 Ivs1 = []int {1, 2, 3}
	i1.Append(1, 6, 7)
	fmt.Println(i1)

	AppendIvs(i1, 2, 2, 2 ,2)
	fmt.Println(i1)

	changeIvs(i1)
	fmt.Println(i1)
}

/*
Slices in Go are passed by reference. To explain your output, let's look closely at what happens in your example in function add:

func add(p []int) {
p = append(p, 24)
fmt.Println(p)
}

When function add is called, a new pointer is created (p) which points to the passed slice. Then, line "p = append(p, 24)" creates a new slice and updates the pointer (p) to point to the slice. When function add exits, variable person still points to the old slice.

If you want to observe a slice change in main method, you can update passed slice inside your add method instead of creating a new one.




 */