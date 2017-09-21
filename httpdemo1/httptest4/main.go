package main

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/mygotest/httpdemo1/httptest4/crons"
	"github.com/mygotest/httpdemo1/httptest4/utils"
	"time"
)

var favoritesProduct = map[string]float32{
	"2952697": 1399, // 显示器
	"2316993": 333,  // 耳机
}

func init() {
	err := utils.ReadConfigJson("./config.json")
	if err != nil {
		panic(err)
	}
}

func main() {
	time.Now()
	crons.JdCron()
	defer crons.JdStopCron()
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalCh:
		fmt.Println("close")
	}

	// single select can block the app
	//select {}
}
