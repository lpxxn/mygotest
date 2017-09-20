package main

import (

	"fmt"

	"os"
	"os/signal"
	"syscall"


	"github.com/robfig/cron"
	f "github.com/mygotest/httpdemo1/httptest4/favorites"
	"github.com/mygotest/httpdemo1/httptest4/utils"
)



var favoritesProduct = map[string]float32{
	"2952697": 1399, // 显示器
	"2316993": 333,  // 耳机
}



func main() {
	err := utils.ReadConfigJson("./config.json")

	if err != nil {
		fmt.Println("read config error: ", err)
		return
	}
	config := utils.AppConfigInstance()
	fmt.Println(config.Favorite_Products)
	c := cron.New()
	c.AddFunc("0 */1 * * * *", func() {
		for pro, price := range favoritesProduct {
			go f.GetPrice(pro, price)
		}

		fmt.Println("Every hour on the half hour")
	})
	fmt.Println("start")
	c.Start()
	defer c.Stop()


	go utils.SendEmail()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalCh:
		fmt.Println("close")
	}

	// single select can block the app
	//select {}
}


