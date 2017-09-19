package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/robfig/cron"
)

type JdPrice []struct {
	Op string `json:"op"`
	M string `json:"m"`
	ID string `json:"id"`
	P string `json:"p"`
}

func main() {
	c := cron.New()
	c.AddFunc("*/1 * * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.Start()
	defer c.Stop()

	resp, err := http.Get("https://p.3.cn/prices/mgets?skuIds=J_2952697")
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
	select {

	}
}
