package main

import (
	"fmt"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported!!!!", http.StatusInternalServerError)
		return
	}

	//w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Write([]byte("abcde\n"))
	f.Flush()

	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-r.Context().Done():
			fmt.Println("client done")
			return

		case <-t.C:
			fmt.Fprintln(w, "hello world ! ")
			//w.Write([]byte("hello world !"))
			f.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
        fmt.Println("listen in: 9100")
	if err := http.ListenAndServe(":9100", nil); err != nil {
		panic(err)
	}
}
