package main

import (
	"time"
	"fmt"
	"github.com/mygotest/timetest/mascot"
)

func main() {
	curTime := time.Now()
	f1 := "2006-01-02 15:04:05"
	parTime, _ := time.Parse(f1, "2018-04-09 00:00:00")
	//parTime, _ := time.Parse(time.RFC3339, "2018-04-09 00:00:00")
	fmt.Println(parTime)
	du := parTime.Sub(curTime)
	fmt.Println("du : ", du)
	if du >= 0 {
		fmt.Println("big")
	}

	du2 := parTime.Sub(parTime)
	fmt.Println(du2)

	du3 := curTime.Sub(parTime)
	fmt.Println("du3 : ", du3)
	d := time.Duration(3)
	fmt.Println(d)
	fmt.Println(mascot.MascotAlpaca)
	fmt.Println(mascot.MascotBuddha)
}