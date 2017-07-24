package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("----", s, "--", i)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("begin receive")
	c <- sum
	fmt.Println("end receive")
}

func main() {
	// 1
	s := []int{1, 3, 2, 10, -1, 5, 0}
	// by default, sends and receives block until the other side is ready.
	// this allows goroutines to synchronize without explicit lock or condition variables
	// asesome
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	// 2
	// buffered channels
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	go say("world")
	say("hello")

}
