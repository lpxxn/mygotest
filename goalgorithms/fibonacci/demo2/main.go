package main

import (
	"fmt"
)

func fibonacci(i int) int{
	if i == 0 {
		return 0
	}
	if i < 2 {
		return 1
	}
	 rev := fibonacci(i - 1) + fibonacci(i - 2)
	 return rev
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
