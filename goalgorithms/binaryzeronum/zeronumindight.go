package main

import (
	"fmt"
	"strconv"
)

func main() {

	count := 0
	fmt.Println("count 1")
	// 1
	x := 29
	fmt.Println(strconv.FormatInt(int64(x), 2))
	for x > 0 {
		count++
		x = x&(x-1)
	}
	fmt.Println(count)


	fmt.Println("count 0")
	// 0
	count = 0
	n := int32(29)
	//n := int64(29)
	fmt.Println("number : ", strconv.FormatInt(int64(n), 2))
	for n > 0 {
		count++
		n = n|(n+1)
	}
	fmt.Println(count)
}
