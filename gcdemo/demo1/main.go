
package main

import (
	"runtime"
	"fmt"
)

func mk2() {
	//b := new([10000]int64)
	//_ = b
	//println(b, "stored at", &b)

	c := make(chan int64, 10000)
	c <- 1
}

func mk1() { mk2() }

func main() {
	RM()
	for i := 0; i < 10; i++ {
		mk1()
		RM()
		runtime.GC()
	}
	RM()
}

func RM(){
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc)
}