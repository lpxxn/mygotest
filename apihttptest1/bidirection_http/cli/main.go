package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
		writer.Write([]byte(" world"))
		writer.Write([]byte("!"))

	})
	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
}
