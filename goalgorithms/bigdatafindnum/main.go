package main

import (
	"fmt"
	"reflect"
	"math"
)

// 40 2亿个无符号型正好是 0~4294967296 42亿多
//
const total = (1 << 32)/8
var total40 = 4 * int(math.Pow(10, 9))
var total40bits = total40/8
var total2 = math.Pow(2, 32)
func main() {
	//bigData := []int {1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(total))
	fmt.Println(total)

	fmt.Println("total40: ", total40, " bits :", total40bits)
	fmt.Println(1<<32)
	fmt.Println(int(total2))
	allBits := make([]byte, total)
	fmt.Println(len(allBits))
}
