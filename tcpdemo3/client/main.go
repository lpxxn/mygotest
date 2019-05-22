package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	//err := conn.(*net.TCPConn).SetKeepAlive(true)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		// Make sure to set keep alive so that the connection doesn't die
		err := tcpConn.SetKeepAlive(true)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = tcpConn.SetKeepAlivePeriod(30 * time.Second)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	notify := make(chan error)

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				notify <- err
				if io.EOF == err {
					close(notify)
					return
				}
			}
			if n > 0 {
				fmt.Println("unexpected data: %s", buf[:n])
			}
		}
	}()

	for {
		select {
		case err := <-notify:
			fmt.Println("connection dropped message", err)
			if err == io.EOF {
				fmt.Println("connection to server was closed")
				return
			}
			break
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1, still alive")
		}
	}

}