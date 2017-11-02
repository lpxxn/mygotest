package main

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/mygotest/httpdemo1/httptest4/crons"
	"github.com/mygotest/httpdemo1/httptest4/utils"
)

func init() {
	err := utils.ReadConfigJson("./config.json")
	if err != nil {
		panic(err)
	}
}

func main() {
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

// container time issue
// docker run -it -d -p 10001:10001 -v /etc/localtime:/etc/localtime:ro  3ee
