package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/hello", controller)

	log.Fatal(http.ListenAndServe(":8090", handler))
}

func controller(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("hello"))
	case <-ctx.Done():
		// If the request gets cancelled, log it
		// to STDERR
		fmt.Fprint(os.Stderr, "request cancelled\n")
	}
}
