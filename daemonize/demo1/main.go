package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// go build
// ./demo1 -d > a.log 2>&1
func main() {
	args := os.Args
	daemon := false
	fmt.Println(args)
	for k, v := range args {
		if v == "-d" {
			daemon = true
			args[k] = ""
		}
	}
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		log.Println(err)
		return
	}

	if daemon {
		Daemonize(args...)
		return
	}

	defer file.Close()
	for {

		file.Write([]byte(strconv.Itoa((int)(time.Now().Unix())) + "\n"))

		time.Sleep(time.Second * 1)
	}
}

func Daemonize(args ...string) {
	var arg []string
	fmt.Println("daemonize ", args)

	if len(args) > 1 {
		arg = args[1:]
	}
	fmt.Println("daemonize2 ", args)
	cmd := exec.Command(args[0], arg...)
	cmd.Env = os.Environ()
	cmd.Start()
}
