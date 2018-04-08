package main

import (
	"fmt"
	"strconv"
	"time"

	"math/rand"
	"os"
	"os/signal"
)

var chs []chan string
func main() {
	for i := 0; i< 10; i++ {
		ch := make(chan string, 100)
		chs = append(chs, ch)
	}
	go func() {
		//for i := 0; i< 100; i++ {
		go func() {
			for {
				//time.Sleep(time.Second/5)
				time.Now()
				v := Random(0, 999999)
				chs[v%10] <- strconv.Itoa(v)
				v = Random(0, 9999)
				chs[v%10] <- strconv.Itoa(v)
			}
		}()

		for {
			v := Random(0, 999999)
			chs[v%10] <- strconv.Itoa(v)
		}
	}()

	//i := 0
	//for i < 10{
	//	str, ok := <- chs[i]
	//	i++
	//	if ok {
	//		fmt.Println(str)
	//	}
	//}
	for i := 0; i< 10; i++ {
		go processF(i)
	}

	fmt.Println("sleep....")
	time.Sleep(time.Second * 5)
	for i := 0; i< 10; i++ {
		//time.Sleep(time.Second * 5)
		chs[i] <- strconv.Itoa(i)
	}


	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, os.Interrupt)
	quit := make(chan bool)

	go func() {
		for _ = range stopSignal {
			fmt.Println("Receive an interrup, Begin Stop.....")

			quit <- true
		}
	}()
	fmt.Println("Running Service ....")
	<-quit
	fmt.Println("Stop Server")
}

func processF(idx int) {
	defer func() {
		if recover()!= nil {
			fmt.Println("error")
		}
	}()
	//for {
	//	str, ok := <- chs[idx]
	//	if ok {
	//		fmt.Println(str)
	//	}
	//}

	for str := range chs[idx] {
		fmt.Println(str)
	}
}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}