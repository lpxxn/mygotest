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
	"github.com/lpxxn/gomail"
	"crypto/tls"
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

/*

 */
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
	c.AddFunc("0 */1 * * * *", func() {
		for pro, price := range favoritesProduct {
			go GetPrice(pro, price)
		}

		fmt.Println("Every hour on the half hour")
	})
	fmt.Println("start")
	c.Start()
	defer c.Stop()



	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalCh:
		fmt.Println("close")
	}

	// single select can block the app
	//select {}
}

func sendEmail() {
	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "p.li@angaomeng.com", "AgmLip123p.li")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}


	m := gomail.NewMessage()
	m.SetHeader("From", "p.li@angaomeng.com")
	m.SetHeader("To", "lpxxn@foxmail.com", "mi_duo@126.com", "p.li@angaomeng.com")
	m.SetHeader("Subject", "Test")
	m.SetBody("text/html", "Hello <b>你好</b> and <i>我是李鹏</i>测试成功!")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("")
}
