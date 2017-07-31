package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:9091")

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Text to send : ")
		text, _ := reader.ReadString('\n')

		// send to socket
		fmt.Fprintf(conn, text+"\n")

		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("message from server : ", message)
	}
}
