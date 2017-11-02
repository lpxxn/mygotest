package main

import (
	"fmt"
	"github.com/mahonia"
	"log"
	"net"
)

// google mahonia
// https://code.google.com/archive/p/mahonia/source/default/source

func main() {
	conn, err := net.Dial("tcp", "MTK-DEV-ELB-MT4-INTERFACE-2021849240.cn-north-1.elb.amazonaws.com.cn:3390")

	if err != nil {
		fmt.Println(err)
	}
	//message := "Wact=symbol_info&uid=7000&sym=GBPJPY&\r\nQUIT\r\n"
	message := "Wact=symbol_info&uid=7000&sym=HLG港敦金\r\nQUIT\r\n"
	//src:="编码转换内容内容"
	//enc := mahonia.NewEncoder("GBK")
	enc := mahonia.NewEncoder("GBK")
	output := enc.ConvertString(message)
	fmt.Println(output)
	// 1
	conn.Write([]byte(output))
	// 2
	// fmt.Fprintf(conn, output)

	// 3

	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)

	_, err = conn.Read(buff)
	if err != nil {
		fmt.Printf("err", err)
	}
	fmt.Printf("Receive: %s \n", buff)
	dec := mahonia.NewDecoder("gbk")
	// convert to utf 8
	if ret, ok := dec.ConvertStringOK(string(buff)); ok {

		fmt.Println("GBK to UTF-8: ", ret, " bytes:")

	}
}
