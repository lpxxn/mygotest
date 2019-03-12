package main

import "fmt"

type Ct1 struct {
	A chan int
}

type Ct2 struct {
	B chan int
}

func main() {
	newCh := make(chan int)

	c1 := Ct1{A: newCh}

	c2 := Ct2{B: c1.A}

	go func() {
		c2.B <- 1
		c1.A <- 2
		close(c2.B)
		// error
		// c1.A <- 3
	}()

	for v := range newCh {
		fmt.Println(v)
	}

}
