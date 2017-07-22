package main

import (
	"fmt"
	"time"
)

type MyFloat64 float64
type MyErr struct {
	When time.Time
	What string
}

func (e *MyErr) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func (f MyFloat64) Error() string {
	return fmt.Sprintf("the float64 error %g", f)
}

func Sqrt(v float64) (float64, error) {
	if v <= 4 {
		return 0, MyFloat64(v)
	}

	return v * 2, nil
}

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

	if v, err := Sqrt(3); err != nil {
		fmt.Println(v, err)
	}
}
