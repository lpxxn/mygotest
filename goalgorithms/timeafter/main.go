package main

import (
	"time"
	"fmt"
)

func getData(url string, ch chan<- bool) {
	time.Sleep(1 * time.Second)
	ch <- true
}

//func IsReachable(urls []string) bool {
//	ch := make(chan bool, len(urls))
//}


func main() {
	ch := make(chan bool)

	go time.AfterFunc(time.Second * 3, func() {
		ch <- false
	})

	rev := <- ch
	fmt.Println(rev)

	doSomething()
	//fmt.Println(doSomething())
}


func doSomething() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		}
		
	}()

	defer func() {
		panic("defer error")
	}()

	fmt.Println("Running...")
	panic("run error")
}

