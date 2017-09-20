package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/lpxxn/gomail"
)


type EmailInfo struct {
	Host string `json:"host"`
	Port int `json:"port"`
	UserName string	`json:"user_name"`
	Pwd string	`json:"pwd"`
	SendTo []string `json:"send_to"`
}

func init() {
}

func SendEmail() {
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
	fmt.Println("send email over")
}