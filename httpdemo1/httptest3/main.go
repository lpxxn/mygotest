package main

import (
	"fmt"
	"io"
	"net/http"
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

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cb", IndexHandler)
	http.HandleFunc("/lp/cb", IndexHandler)
	if err := http.ListenAndServe(":9100", nil); err != nil {
		panic(err)
	}
}
