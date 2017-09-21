package crons

import (
	"github.com/robfig/cron"
	"fmt"
	"github.com/mygotest/httpdemo1/httptest4/utils"
	"github.com/mygotest/httpdemo1/httptest4/bll"
)
var jd_cron = cron.New()

func JdCron() {
	jd_cron.AddFunc("0 */1 * * * *", func() {
		config := utils.AppConfigInstance()
		bll.GetPrice(&config.JdProductInfo)
		fmt.Println("Every minute ")
	})
	jd_cron.Start()
	fmt.Println("start jd cron")
}

func JdStopCron() {
	jd_cron.Stop()
}
