package crons

import (
	"fmt"
	"github.com/mygotest/httpdemo1/httptest4/bll"
	"github.com/mygotest/httpdemo1/httptest4/utils"
	"github.com/robfig/cron"
	"time"
)

var jd_cron = cron.New()

func JdCron() {
	jd_cron.AddFunc("0 */1 * * * *", func() {
		config := utils.AppConfigInstance()
		bll.GetPrice(&config.JdProductInfo)
		fmt.Println("Every minute  current time", time.Now())
	})
	jd_cron.Start()
	fmt.Println("start jd cron")
}

func JdStopCron() {
	jd_cron.Stop()
}
