package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
	"testing"
)

// https://godoc.org/net/http/httputil#DumpRequest

func TestDumpResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		// %q 不会格式化输出
		fmt.Fprintf(w, "%q\n", dump)
		fmt.Fprintf(w, "-----------\n")
		fmt.Fprintf(w, "%s\n", string(dump))
		fmt.Fprintf(w, "-----------\n")
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, " body: %s \n header: %#v", string(body), r.Header)
	}))
	defer ts.Close()

	const body = `{ "desc": "Go is a general-purpose language designed with systems programming in mind.", "age": 10}`
	req, err := http.NewRequest("POST", ts.URL, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("test", "abcdef")
	req.Header.Add("client_id", "aaaaa")
	req.Header.Add("", "")
	req.Host = "www.example.org"
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}
