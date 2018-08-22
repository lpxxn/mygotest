package main

import (
	"time"
	"fmt"
	"os"
	"os/signal"
)


type TChanTest  struct {
	Id string
	Stop chan bool
}

func (t *TChanTest) Run() {
	t1 := time.Tick(time.Second * 2)

	t2 := time.Tick(time.Second)

	for {
		select {
		case <-t1:
			fmt.Println("t1 start running.....")

		case <- t2:
			fmt.Println("t2....................")

		case <-t.Stop:
			fmt.Println("service stopping.....")
			return
		}
	}
}


func main() {
	m1 := make(map[string]*TChanTest)

	id := "abcde"
	m1[id] = &TChanTest{Id: id, Stop: make(chan bool)}

	go m1[id].Run()


	time.AfterFunc(time.Second * 11, func() {
		m1[id].Stop <- true
		fmt.Println("stop <- true")
	})

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("server running....")
	<- stop
}
