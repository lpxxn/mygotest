package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(a []int, c chan int) {
	fmt.Println("summing: ", a)
	total := 0
	for _, v := range a {
		total += v
	}
	//fmt.Println("send to c",total)
	c <- total  // send total to c
}
func main() {
	runtime.GOMAXPROCS(1)

	a := []int{7, 2, 8,134,23,23,1,23,1234,143, -9, 4, 0, 1234}

	c := make(chan int)

	for i := 0; i < 100; i++ {
		go sum(a[:len(a)/2], c)
		go sum(a[len(a)/2:], c)

		go sum(a[:len(a)/2], c)
		go sum(a[len(a)/2:], c)

		//var x = 0
		x := <-c
		fmt.Println(x)
		x = <-c
		//fmt.Println(x, "-------------")
	}
	time.Sleep(time.Second * 5)
}
