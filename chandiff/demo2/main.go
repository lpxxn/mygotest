package main

import (
	"fmt"
	"sync"
)

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
	wg := sync.WaitGroup{}
	wg.Add(2)
	go getValue(c1.A, 1, &wg)

	byteChan := make(chan []byte)

	go func() {
		c2.B <- 1
		c1.A <- 2
		close(c2.B)
		// error
		// c1.A <- 3
		getValue(c1.A, 2, &wg)
	}()
	go func() {
		for v := range 	byteChan {
			_ = v
		}
	}()
	byteValue := []byte("abcdef")
	byteChan <- byteValue
	for v := range newCh {
		fmt.Println("main func get value: ", v)
	}
	byteChan <- byteValue

	if v, ok := <-newCh; ok {
		fmt.Println("newCh value", v)
	}
	if _, ok := <-c1.A; !ok {
		fmt.Println("C1: A is closed")
	}
	fmt.Println("byteValue:", string(byteValue))

	byteValue = nil
	byteChan <- byteValue

	fmt.Println("wait--------")

	wg.Wait()
	fmt.Println("byteValue:", string(byteValue))
}

func getValue(ch1 <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	v, ok := <-ch1
	fmt.Println("id: ", id)
	if !ok {
		fmt.Println("not ok already closed, v:", v)
		return
	}
	fmt.Println("getValue value: ", v)

}


