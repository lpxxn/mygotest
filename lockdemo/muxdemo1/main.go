package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type M1 struct {
	mux sync.Mutex
	v   int
}

// must be pointer because all parameters are passed by value
func (m *M1) Func1() {
	m.mux.Lock()
	defer m.mux.Unlock()
	fmt.Println("func1 running")
	m.v++
	time.Sleep(time.Second * 3)
	fmt.Println("111 end")
}

func (m *M1) Func2() {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.v++
	fmt.Println("2222 running")
	time.Sleep(time.Second * 3)
	fmt.Println("222 end")
}

func main() {
	m1 := M1{}

	go m1.Func2()
	go m1.Func1()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	fmt.Println("stop server....")
}
