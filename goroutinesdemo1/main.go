package main

import (
	"fmt"
	"time"
)

var value int = 1

func say(s string) {

	for i := 0; i < 5; i++ {
		value += 1
		time.Sleep(500 * time.Millisecond)

		fmt.Println("----", s, "--", i)

	}
}

func main() {
	go say("world")
	go say("world2")
	go say("world3")
	say("hello")
}
