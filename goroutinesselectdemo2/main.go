package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(10 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")

		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(30 * time.Millisecond)
		}
	}
}
