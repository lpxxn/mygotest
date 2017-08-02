package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mahonia"
	"io/ioutil"
	"net"
	"os"
)

type ConfigObject struct {
	Service string `json: "service"`
	Uuid    string `json: "uuid"`
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	if l := len(args); l < 2 {
		fmt.Println("please input act and parameters. ex: ")
		fmt.Println("testtcpmtinterface.exe symbol_info sym=HLG港敦金")
		return
	}
	configObj, fileErr := ConfigObjByJson()
	if fileErr != nil {
		fmt.Println("read json error: ", fileErr)
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString("Wact=")
	buffer.WriteString(args[0])
	buffer.WriteString("&uid=")
	buffer.WriteString(configObj.Uuid)
	params := args[1:]
	for _, item := range params {
		buffer.WriteString("&" + item)
	}
	buffer.WriteString("\r\nQUIT\r\n")

	enc := mahonia.NewEncoder("GBK")
	bufStr := buffer.String()
	message := enc.ConvertString(bufStr)
	fmt.Printf("utf-8 msg: %s \n", bufStr)
	//fmt.Println("BGK msg: ", message)

	fmt.Println("service address :", configObj.Service)
	conn, err := net.Dial("tcp", configObj.Service)
	if err != nil {
		fmt.Println("connect server error :", err)
		return
	}

	conn.Write([]byte(message))

	revBuff := make([]byte, 1024)
	_, err = conn.Read(revBuff)

	if err != nil {
		fmt.Printf("receive msg error : %s \n", err)
		return
	}
	dec := mahonia.NewDecoder("GBK")
	ret := dec.ConvertString(string(revBuff))
	fmt.Println("--------------\n\n")
	fmt.Printf("receive msg : %s \n", ret)
}

func ConfigObjByJson() (*ConfigObject, error) {
	file, fileErr := ioutil.ReadFile("./config.json")
	if fileErr != nil {
		return nil, fileErr
	}
	var jsonObj *ConfigObject = &ConfigObject{}
	json.Unmarshal(file, jsonObj)
	return jsonObj, nil
}
