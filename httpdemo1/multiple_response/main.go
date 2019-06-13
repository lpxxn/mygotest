package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported!!!!", http.StatusInternalServerError)
			return
		}
		notify := r.Context().Done()
		rollMsg := make(chan string, 5)
		defer func() {
			close(rollMsg)
		}()

		w.Header().Add("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		go func() {
			ticker := time.NewTicker(time.Second / 2)
			defer ticker.Stop()
			for {
				select {
				case <-notify:
					fmt.Println("client closed stop write roll msg")
					// client is close
					return
				case t := <-ticker.C:
					rollMsg <- t.String()
				}

			}
		}()
		for {
			select {
			case <-notify:
				fmt.Println("client is gone, stop write data to client")
				// client is close
				return
			case msg := <-rollMsg:
				if _, err := w.Write([]byte(msg)); err != nil {
					return
				}
				f.Flush()

			}
		}
	})

	if err := http.ListenAndServe(":9111", nil); err != nil {
		panic(err)
	}
}
