package main

import (
	"crypto/tls"

	// go get gopkg.in/gomail.v2
	// donot support dep command
	"gopkg.in/gomail.v2"
)

func main() {
	d := gomail.NewDialer("smtp.mxhichina.com", 465, "lipeng@rrzhuan.com", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "lipeng@rrzhuan.com")
	m.SetHeader("To", "lpxxn@foxmail.com", "mi_duo@126.com", "zhangyanan@rrzhuan.com", "yejiani@rrzhuan.com")
	m.SetHeader("Subject", "Test")
	m.SetBody("text/html", "Hello <b>你好</b> and <i>我是李鹏</i>测试成功!")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
