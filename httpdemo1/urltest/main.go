package main

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"net/url"
)

type YpCbInfo []struct {
	Sid             int64  `json:"sid"`
	UID             string `json:"uid"`
	UserReceiveTime string `json:"user_receive_time"`
	ErrorMsg        string `json:"error_msg"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
}

func main() {
	//fmt.Println(url.QueryEscape("+41545"))
	//fmt.Println(url.QueryEscape("123456798"))

	m := mahonia.NewDecoder("UTF-8")

	str := "sms_status=%255B%257B%2522sid%2522%253A23839436406%252C%2522uid%2522%253Anull%252C%2522user_receive_time%2522%253A%25222018-05-03%2B09%253A19%253A45%2522%252C%2522error_msg%2522%253A%2522DELIVRD%2522%252C%2522mobile%2522%253A%252215117950565%2522%252C%2522report_status%2522%253A%2522SUCCESS%2522%257D%255D"

	un, e := url.QueryUnescape(str)
	fmt.Println(un, e)

	un, e = url.QueryUnescape(un)
	fmt.Println(un, e)
	m.Translate([]byte(str), true)

	p, e := url.ParseQuery(un)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(p)
	cb_str := p.Get("sms_status")
	fmt.Println(cb_str)
	cb_info := make(YpCbInfo, 0)

	json.Unmarshal([]byte(cb_str), &cb_info)
	fmt.Println(cb_info)

	//str2 := `{"level":"info","ts":1525260979.8594625,"msg":"yp cb :sms_status=%255B%257B%2522sid%2522%253A23832105116%252C%2522uid%2522%253Anull%252C%2522user_receive_time%2522%253A%25222018-05-02%2B19%253A36%253A18%2522%252C%2522error_msg%2522%253A%2522DELIVRD%2522%252C%2522mobile%2522%253A%252215117950565%2522%252C%2522report_status%2522%253A%2522SUCCESS%2522%257D%255D","curenttime":"2018-05-02T19:36:19+08:00"}`
	//un3, err := url.QueryUnescape(str2)
	//un3, err = url.QueryUnescape(un3)

	//fmt.Println(un3, err)
	//u, err := url.ParseQuery(str)
	//fmt.Println(err, u)
}
