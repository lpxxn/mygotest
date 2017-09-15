package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 3)
	fmt.Println("len: ", len(slice1), "cap: ", cap(slice1), "array :", slice1)
	slice1 = append(slice1, 1)
	fmt.Println("len: ", len(slice1), "cap: ", cap(slice1), "array :", slice1)

	slice2 := make([]int, 3, 7)
	fmt.Println("len: ", len(slice2), "cap: ", cap(slice2), "array :", slice2)

	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("len: ", len(slice3), "cap: ", cap(slice3), "array :", slice3)

	slice4 := []int{5: 0}
	fmt.Println("len: ", len(slice4), "cap: ", cap(slice4), "array :", slice4)

	var slice5 []int
	fmt.Println("len: ", len(slice5), "cap: ", cap(slice5), "array :", slice5)
	slice5 = append(slice5, 4)
	fmt.Println("len: ", len(slice5), "cap: ", cap(slice5), "array :", slice5)

	slice6 := []int{}
	fmt.Println("len: ", len(slice6), "cap: ", cap(slice6), "array :", slice6)
	slice6 = append(slice6, 2)
	fmt.Println("len: ", len(slice6), "cap: ", cap(slice6), "array :", slice6)

	slice7 := make([]int, 0)
	fmt.Println("len: ", len(slice7), "cap: ", cap(slice7), "array :", slice7)
	slice7 = append(slice7, 7)
	fmt.Println("len: ", len(slice7), "cap: ", cap(slice7), "array :", slice7)
}
