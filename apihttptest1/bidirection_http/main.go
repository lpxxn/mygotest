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
	//f, ok := w.(http.Flusher)
	//if !ok {
	//	http.Error(w, "streaming unsupported!!!!", http.StatusInternalServerError)
	//	return
	//}
	btys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body err: ", err)
		w.Write([]byte("read body err" + err.Error()))
		return
	}
	fmt.Println("received msg: ", string(btys))
	rsp := http.Response{
		Status: http.StatusText(200),
		Request:r,
		Body: ioutil.NopCloser(strings.NewReader("sfdsdfsdf\n")),
	}
	time.AfterFunc(time.Second, func() {
		rsp.Write(w)
		fmt.Fprint(w, "adfasaaaadf\n")
		//f.Flush()
	})

	w.Write([]byte("hello world ! "))
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
	println("print response 1---")
	println(string(rspBody))
	var r2 *http.Request

	resp2, err := http.ReadResponse(bufio.NewReader(conn), r2)
	defer resp2.Body.Close()
	rspBody2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	println("print response 2---")
	println(string(rspBody2))

	var r3 *http.Request
	resp3, err := http.ReadResponse(bufio.NewReader(conn), r3)
	defer resp3.Body.Close()
	rspBody3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	println("print response 3---")
	println(string(rspBody3))
	//http.DefaultClient.Do(r)
}
