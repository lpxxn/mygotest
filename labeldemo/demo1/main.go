package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go Fun1(c1, c2, c3)
	time.Sleep(5 * time.Second)
	MyLoop:
	for {
		select {
		case <- c1:
			fmt.Println("c1")
		case <-c2:
			fmt.Println("c2")
		case <-c3:
			fmt.Println("c3")
			break
		default:
			fmt.Println("MyLoop")
			break MyLoop
		}
	}
}
func Fun1(c1, c2, c3 chan<- int) {
	c1 <- 10
	time.AfterFunc(10, func() {
		c3<-33
		go Fun1(c1, c2, c3)
	})

	c2 <- 3

}