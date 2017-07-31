package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "MTK-DEV-ELB-MT4-INTERFACE-2021849240.cn-north-1.elb.amazonaws.com.cn:3390")

	if err != nil {
		fmt.Println(err)
	}
	message := "Wact=symbol_info&uid=7000&sym=HKG港金\r\nQUIT\r\n"
	conn.Write([]byte(message))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)

	_, err = conn.Read(buff)
	if err != nil {
		fmt.Printf("err", err)
	}
	fmt.Printf("Receive: %s \n", buff)

}
