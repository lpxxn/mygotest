package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server ....")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", "127.0.0.1:9091")

	// accept connection on port
	conn, _ := ln.Accept()
	// run loop forever (or util ctrl-c)
	for {
		// will listen for message to process ending in newline(\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Println("Message Received: ", string(message))

		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
