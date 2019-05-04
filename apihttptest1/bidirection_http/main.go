package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.Host)
	//fmt.Fprintln(w, "hello world")
	w.Write([]byte("hello "))
	w.Write([]byte("world !"))
	w.Write([]byte(io.EOF.Error()))
}

func bidirectionFunc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	btys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body err: ", err)
		w.Write([]byte("read body err" + err.Error()))
		return
	}
	fmt.Println("received msg: ", string(btys))
	w.Write(append([]byte("i received "), btys...))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cb", IndexHandler)
	http.HandleFunc("/bd", bidirectionFunc)
	time.AfterFunc(time.Second, CbRequest)
	if err := http.ListenAndServe(":9100", nil); err != nil {
		panic(err)
	}
}

var testUrl = "http://127.0.0.1:9100/bd?name=aaaaaaa"
func CbRequest() {
	uri, _ := url.Parse(testUrl)
	body1 := `{"name": "li"}`
	r, _ := http.NewRequest(http.MethodPost, uri.String(), strings.NewReader(body1))
	_ = net.JoinHostPort(uri.Host, uri.Port())
	conn, err := net.Dial("tcp", "127.0.0.1:9100")
	if err != nil {
		panic(err)
	}
	if err := r.Write(conn); err != nil {
		fmt.Println(err)
		return
	}
	b2 := `{"name": "peng"}`
	r.ContentLength = int64(len(b2))
	r.Body = ioutil.NopCloser(strings.NewReader(b2))
	if err := r.Write(conn); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("begin read:============")
	resp, err := http.ReadResponse(bufio.NewReader(conn), r)
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(string(rspBody))
	resp2, err := http.ReadResponse(bufio.NewReader(conn), r)
	defer resp2.Body.Close()
	rspBody2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(string(rspBody2))
	//http.DefaultClient.Do(r)
}