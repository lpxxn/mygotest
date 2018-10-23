package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)
func main() {
	cron1 := cron.New()
	defer cron1.Stop()

	// 秒
	cron1.AddFunc("*/10 * * * * ?", func() {
		//fmt.Println("second")
		})

	cron1.AddFunc("0 */30 * * * ?", func() {
		fmt.Println("min 30 ", time.Now(), " hour : ", time.Now().Hour())
		if time.Now().Hour() == 0 {
			fmt.Println("time is zero")
		}
	})

	cron1.AddFunc("0 */1 * * * ?", func() {
		fmt.Println("min 1 ", time.Now())
	})
	cron1.AddFunc("0 0 * * * ?", func() {
		fmt.Println("hour 1 ", time.Now(), " hour: ", time.Now().Hour())
	})
	
	// 晚上0点
	cron1.AddFunc("0 0 0 * * ?", func() {
		fmt.Println("zero")

		fmt.Println(time.Now())
		fmt.Println(time.Now().AddDate(0, 0, -1))
	})
	fmt.Println("begin ", time.Now())

	cron1.Start()
	select{}
}
