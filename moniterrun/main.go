package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"sync"

	"github.com/creack/pty"
)

var exitCh = make(chan os.Signal)

func main() {
	cmd := exec.Command("./webserverdemo1", "-env=local -a=bcd", "-abc=ddd")
	//cmd := exec.Command("./basi", "-env=local")
	cmd.Env = append(cmd.Env, "ENV=dev")
	//cmd.Stdout = os.Stdout
	fmt.Println(cmd.String())
	fmt.Println(cmd.Env)
	//cmdReader, err := cmd.StdoutPipe()
	//if err != nil {
	//	panic(err)
	//}
	//scanner := bufio.NewScanner(cmdReader)
	//go func() {
	//	for scanner.Scan() {
	//		fmt.Println(scanner.Text())
	//		//fmt.Println(scanner.Text())
	//	}
	//}()
	//
	//if err := cmd.Start(); err != nil {
	//	panic(err)
	//}
	//
	//if err := cmd.Wait(); err != nil {
	//	panic(err)
	//}
	f, err := pty.Start(cmd)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, f)

	signal.Notify(exitCh, os.Interrupt)
	fmt.Println("run service")
	<-exitCh
	fmt.Println("exit service")
}

func RunServer() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		cmd := exec.Command("webserverdemo1")
		cmd.Env = append(cmd.Env, "ENV=dev")
		cmd.Run()
	}()
	go func() {

		select {
		case <-exitCh:
		}
	}()
}
