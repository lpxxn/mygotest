package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"compress/gzip"
)

func main() {
	client := new(http.Client)
	url := "http://13.124.215.155:9002/act/wlgroups"
	payload := strings.NewReader(`{"uid": 1001}`)
	request, _ := http.NewRequest("POST", url, payload)
	request.Header.Add("accept-encoding", "gzip")
	request.Header.Add("content-type", "application/json")

	response, _ := client.Do(request)
	defer response.Body.Close()

	//body, _ := ioutil.ReadAll(response.Body)
	//body2, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(body2)

	//fmt.Println(string(body))
	//enc := mahonia.NewEncoder("GBK")
	//message := enc.ConvertString(string(body))
	//fmt.Println(message)

	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(response.Body)
		fmt.Println("content-encoding is gzip")
	default:
		reader = response.Body
		fmt.Println("content-encoding not gzip")
	}
    defer reader.Close()
	body2, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(body2))
	enc := mahonia.NewEncoder("GBK")
	message := enc.ConvertString(string(body2))
	fmt.Println(message)
}
