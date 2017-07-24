package main

import (
	"time"
	"fmt"
)

type MyErr struct {
	When time.Time
	What string
}

func(e *MyErr) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}


type MyFloat64 float64


func run() error {
	return &MyErr{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
