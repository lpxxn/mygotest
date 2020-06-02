package main

import (
	"fmt"

	"github.com/mygotest/backoff/demo1/backoff"
)

func main() {
	e := backoff.DefaultExponential
	for i := 1; i < 200; i++ {
		t := e.Backoff(i)
		fmt.Printf("retires: %d, duration: %f \n", i, t.Seconds())
	}
}
