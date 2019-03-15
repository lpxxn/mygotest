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
	go m1.Func2()
	go m1.Func1()

	go m1.Func1()



	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	fmt.Println("stop server....")
}
/*
1. 读锁不能阻塞读锁，引入readerCount实现

2. 读锁需要阻塞写锁，直到所以读锁都释放，引入readerSem实现

3. 写锁需要阻塞读锁，直到所以写锁都释放，引入wirterSem实现

4. 写锁需要阻塞写锁，引入Metux实现
 */