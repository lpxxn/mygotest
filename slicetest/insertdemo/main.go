package main

import (
	"fmt"
)

func main() {
	var a []int = []int{1, 2, 3, 4}
	fmt.Println(a)

	//b := []int {4: 5, 6}
	b := []int {5, 6}

	insertIndex := 2

	tmp := append([]int{}, a[insertIndex:]...)
	fmt.Println(tmp)
	fmt.Println("a[0:index]", a[0:insertIndex])
	a = append(a[0:insertIndex], b...)
	fmt.Println(a)
	a = append(a, tmp...)

	//copy(b[1:], a)
	//a = b
	//a = append(b, a...)
	fmt.Println(a)
	// delete

	delIndex := 3
	a = append(a[:delIndex], a[delIndex+1 :]...)
	fmt.Println(a)
}

