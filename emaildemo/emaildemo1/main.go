package main

import (
	"crypto/tls"

	// go get gopkg.in/gomail.v2
	// donot support dep command
	"gopkg.in/gomail.v2"
)

func main() {
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

}
