package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mahonia"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

type msg struct {
	Num     int
	Command string
	OutPut  string
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go echo(conn)
}

func echo(conn *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
	}()
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)
		cmd := exec.Command("cmd", "/c", m.Command)
		var buf bytes.Buffer
		cmd.Stdout = &buf
		cmd.Start()

		// Use a channel to signal completion so we can use a select statement
		done := make(chan error)
		go func() { done <- cmd.Wait() }()

		// Start a timer
		timeout := time.After(10 * time.Second)

		// The select statement allows us to execute based on which channel
		// we get a message from first.
		select {
		case <-timeout:
			// Timeout happened first, kill the process and print a message.
			cmd.Process.Kill()
			fmt.Println("Command timed out")
			m.OutPut = "time out"
		case err := <-done:
			// Command completed before timeout. Print output and error if it exists.
			fmt.Println("Output:", buf.String())

			if err != nil {
				m.OutPut = fmt.Sprintf("%v", err)
			} else {
				dec := mahonia.NewDecoder("GBK")
				//dec := mahonia.NewDecoder("UTF-8")
				ret := dec.ConvertString(buf.String())
				fmt.Println(ret)
				m.OutPut = ret
			}

		}
		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
		//cmd, err := exec.Command("cmd", "/c", m.Command).Output()

	}
}
