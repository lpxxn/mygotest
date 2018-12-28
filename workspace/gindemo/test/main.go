package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int, len(s))
	sum := 0

	go func() {
		for _, v := range s {
			getSum2(v, result)
		}
		close(result)
	}()

	for val := range result {
		sum += val
		fmt.Println("total item val:", sum)
	}
	fmt.Println(sum)
}

func getSum2(val int, sum chan<- int) {
	sum <- int(math.Pow(float64(val), 2))
}
