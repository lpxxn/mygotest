package main

import (
	"runtime"
	"sync"
	"fmt"
	"unsafe"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}

	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numGoroutines = 1000
	wg.Add(numGoroutines)

	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3f kb \n", float64(after - before)/ numGoroutines / 1000)
	done1 := make(chan struct{})
	go func() {done1 <- struct {}{}}()
	//close(done1)
	if _, okc := <-done1; !okc {
		fmt.Println("close")
	}
	v1 := unsafe.Sizeof(done1)

	done2 := make(chan interface{})
	v2 := unsafe.Sizeof(done2)
	fmt.Println("struct :", v1, " interface :", v2)

	var s struct{}

	v1 = unsafe.Sizeof(s)
	var i interface{}
	v2 = unsafe.Sizeof(i)
	fmt.Println("struct :", v1, " interface :", v2)
}







