package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron"
)

type JdPrice []struct {
	Op string `json:"op"`
	M  string `json:"m"`
	ID string `json:"id"`
	P  string `json:"p"`
}

const (
	priceUrl = "https://p.3.cn/prices/mgets?skuIds=J_"
)

var favoritesProduct = map[string]float64{
	"2952697": 1399, // 显示器
	"2316993": 333,  // 耳机
}

///
/// test
func GetPrice(product string, myPrice float64) {
	//2316993  2316993
	//resp, err := http.Get("https://p.3.cn/prices/mgets?skuIds=J_2316993")
	resp, err := http.Get(priceUrl + product)
	if err == nil {
		msg, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			strBody := string(msg)
			fmt.Println("msg: ", strBody)
			var jd JdPrice
			json.Unmarshal(msg, &jd)
			fmt.Println(jd)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func main() {
	c := cron.New()
	c.AddFunc("*/1 * * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.Start()
	defer c.Stop()

	for pro, price := range favoritesProduct {
		go GetPrice(pro, price)
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalCh:
		fmt.Println("close")
	}

	// single select can block the app
	//select {}
}
