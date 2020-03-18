package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/mygotest/timetest/demo2/lib"

	"github.com/mygotest/timetest/mascot"
)

func main() {
	lib.TimeNow()
	//os.Setenv("TZ", "Asia/Shanghai")
	loc, err := time.LoadLocation("America/Atka")
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Local = loc
	lib.TimeNow()

	loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Local = loc
	lib.TimeNow()
	//os.Setenv("TZ", "America/Los_Angeles")
	//fmt.Println(time.Local, "time now: ", time.Now())
	//fmt.Println(time.Local, "time now: ", time.Now())

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
	fmt.Println("du3 : ", du3, "parTime: ", parTime, "is great :", du3 > 0)
	du4 := parTime.Sub(curTime)
	fmt.Println("du4 : ", du4, "parTime: ", parTime, "is great :", du4 > 0)
	d := time.Duration(3)
	fmt.Println(d)
	//fmt.Println(mascot.MascotAlpaca)
	//fmt.Println(mascot.MascotBuddha)
	color.Green(mascot.MascotAlpaca)
	color.Yellow(mascot.MascotBuddha)
	//fmt.Println(mascot.MascotRabbit)
	color.Blue(mascot.MascotRabbit)
}
