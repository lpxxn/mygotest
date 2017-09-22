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

func SendEmail(body  <-chan string) {
	strBody := <- body
	config := AppConfigInstance().EmailInfoConfig
	// 	"smtp.exmail.qq.com",
	d := gomail.NewDialer(config.Host, config.Port, config.UserName, config.Pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", config.UserName)
	m.SetHeader("To", config.SendTo...)
	m.SetHeader("Subject", "降价啦")
	m.SetBody("text/html", strBody)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("send email over", strBody)
}