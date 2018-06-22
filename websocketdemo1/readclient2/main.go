package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"math/rand"
	"time"
	"github.com/gorilla/websocket"
	"net/url"
	"fmt"
)

//var addr = flag.String("addr", "localhost:8080", "http service address")
//var addr = flag.String("addr", "localhost:8101", "http service address")
var addr = flag.String("addr", "192.168.3.34:8101", "http service address")
func main() {
	flag.Parse()
	log.SetFlags(0)
	rand.Seed(time.Now().UnixNano())
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	for i:=0; i < 100; i++ {
		go testclients()
	}

	<-interrupt

	//select {
	//
	//}

}

func testclients() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	var read_count  int64 = 0
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			read_count++
			log.Printf("recv: %s read: %d", message, read_count)

		}
	}()
	uid := rand.Intn(1000)

	err = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"type":"login", "client_name": "auto","uid":"%d","room_id":"111"}`, uid)))
	if err != nil {
		log.Println("write:", err)
		return
	}


	ticker := time.NewTicker(time.Second * 30)


	defer ticker.Stop()
	send_data := []byte(`{"type":"ping", "uid":"110","room_id":"111", "content": "abcdef"}`)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, send_data)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}