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
	cron1.AddFunc("*/1 * * * * *", func() {
		fmt.Println("second")
		})

	cron1.AddFunc("0 */30 * * * ?", func() {
		fmt.Println("min 30")
	})

	// 晚上0点
	cron1.AddFunc("0 0 0 * * ?", func() {
		fmt.Println("zero")

		fmt.Println(time.Now())
		fmt.Println(time.Now().AddDate(0, 0, -1))
	})
	fmt.Println("begin")

	cron1.Start()
	select{}
}
