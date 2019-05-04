package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world ! ")
}

func cbHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	btys, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body err: ", err)
		w.Write([]byte("read body err" + err.Error()))
		return
	}
	fmt.Fprintln(w, "received msg: ", string(btys))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cb", cbHandler)
	if err := http.ListenAndServe(":9100", nil); err != nil {
		panic(err)
	}
}
