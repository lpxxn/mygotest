package main

import (
	"fmt"
	"time"
	"runtime"
)


func RM(){
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m)
}

func main() {
	fmt.Println("begin")
	//runtime.GC() // get up-to-date statistics
	Testgorouting();
	time.Sleep(time.Second * 5)

	fmt.Println("end")
	//time.AfterFunc(time.Second * 5, func() {
	//
	//})
}

func Testgorouting() {

	fmt.Println("test gorouting")
	//RM()
	time.Sleep(time.Second * 3)
	c1 := make(chan int, 20)
	for i := 0; i < 5; i++ {
		c1 <- i
	}
	//RM()
	time.Sleep(time.Second * 3)
	runtime.GC() // get up-to-date statistics

	//RM()
}



