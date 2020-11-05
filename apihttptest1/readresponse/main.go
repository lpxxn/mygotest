package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type MyConnType struct {
}

func (mct *MyConnType) SendMsg(urlStr, msg string) (*http.Response, error) {
	uri, err := url.Parse(urlStr)
	conn, err := net.Dial("tcp", uri.Host+":80")
	if err != nil {
		return nil, err
	}
	fmt.Println(conn.RemoteAddr().String())
	reqB := bytes.NewBufferString(msg)
	defer reqB.Reset()

	req, err := http.NewRequest("GET", urlStr, reqB)
	if err != nil {
		return nil, err
	}
	if err = req.Write(conn); err != nil {
		return nil, err
	}
	//buff := bufio.NewReader(conn)
	resp, err := http.ReadResponse(bufio.NewReader(conn), req)
	return resp, nil
}

func (mct *MyConnType) SendMsg2(urlStr, msg string) (*http.Response, error) {
	uri, err := url.Parse(urlStr)
	conn, err := net.Dial("tcp", uri.Host+":80")
	if err != nil {
		return nil, err
	}
	fmt.Println(conn.RemoteAddr().String())
	//reqB := bytes.NewBufferString(msg)
	//defer reqB.Reset()
	//rc := ioutil.NopCloser(reqB)
	rc := ioutil.NopCloser(strings.NewReader(msg))
	fmt.Println("len of body:", len(msg))
	req := &http.Request{
		Method: "GET",
		URL: uri,
		Body:          rc,
		ContentLength: int64(len(msg)),
		Host:          uri.Host,
	}
	if err = req.Write(conn); err != nil {
		return nil, err
	}
	//buff := bufio.NewReader(conn)
	resp, err := http.ReadResponse(bufio.NewReader(conn), req)
	return resp, nil
}

func (mct *MyConnType) SendMsgFun(urlStr string, f func(urlStr, msg string) (*http.Response, error)) {
	resp, err := f(urlStr, "一二三")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	rev := new(RevData)
	err = json.Unmarshal(body, rev)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("rev data: ", rev)
}

func main() {
	urlStr := "http://my-json-server.typicode.com/lpxxn/mygotest/posts/1"
	mct := new(MyConnType)
	mct.SendMsgFun(urlStr, mct.SendMsg)
	fmt.Println("------------")
	mct.SendMsgFun(urlStr, mct.SendMsg2)

	uri, _ := url.Parse(urlStr)
	if ips, err := net.LookupIP(uri.Host); err == nil {
		for _, ipItem := range ips {
			fmt.Println(ipItem.String())
		}
	}
}

type RevData struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
