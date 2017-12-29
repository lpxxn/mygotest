package main

import (
	"time"
	"fmt"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	//timer1.Stop()

	<-timer1.C
	fmt.Println("timer 1 expired")
	time2 := time.NewTimer(time.Second * 2)

	go func() {
		<- time2.C

		fmt.Println(" timer 2 expired")
	}()








	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	select {

	}
}
