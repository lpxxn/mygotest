package main

import "fmt"

func func1(ch chan bool) {
	for i := 0; i < 5; i++ {
		select {
		case <-ch:
			fmt.Println("Read")
		case ch <- true:
			fmt.Println("write")
		default:
			fmt.Println("Neither")
		}
	}
}

func main() {
	ch1 := make(chan bool)

	func1(ch1)

	fmt.Println("-----------------")

	ch2 := make(chan bool, 1)
	func1(ch2)
}
