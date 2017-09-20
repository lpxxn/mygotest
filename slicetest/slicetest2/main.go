package main

import (
	"fmt"
)

func main() {
	// 创建一个容量和长度均为6的slice
	slice1 := []int{5, 23, 10, 2, 61, 33}

	for index, value := range slice1 {
		fmt.Println("index: ", index, " value: ", value)
	}

	// 可以只用我们关心的元素
	// 只关心value
	for _, value := range slice1 {
		fmt.Println("value ", value)
	}

	// 只关心index
	for index, _ := range slice1 {
		fmt.Println("index: ", index)
	}

	// 对slices1进行切片，长度为2容量为3
	slice2 := slice1[1:3:3]
	fmt.Println("cap", cap(slice2))
	fmt.Println("slice2", slice2)

	//修改一个共同指向的元素
	//两个slice的值都会修改
	slice2[0] = 11111
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	// 增加一个元素
	slice2 = append(slice2, 55555)

	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)

	fmt.Println("\r\n-----")
	slice3 := []int{1, 2, 3}
	fmt.Println("slice2 cap", cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println("slice2 cap", cap(slice3))

}
