package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type M1 struct {
	rw sync.RWMutex
}

func (m *M1) Func1() {
	m.rw.Lock()
	defer m.rw.Unlock()
	fmt.Println("func1 running")
	time.Sleep(time.Second * 3)
	fmt.Println("111 end")
}

func (m *M1) Func2() {
	m.rw.RLock()
	defer m.rw.RUnlock()
	fmt.Println("2222 running")
	time.Sleep(time.Second * 3)
	fmt.Println("222 end")
}

func main() {
	m1 := M1{}

	go m1.Func2()
	go m1.Func1()

	go m1.Func1()



	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	fmt.Println("stop server....")
}
