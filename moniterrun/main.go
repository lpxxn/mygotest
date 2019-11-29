package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"time"

	"github.com/creack/pty"
)

func main() {
	var exitCh = make(chan os.Signal)
	ctx, cancel := context.WithCancel(context.Background())

	signal.Notify(exitCh, os.Interrupt)

	go RunServer(ctx)

	fmt.Println("run service")
	<-exitCh
	cancel()
	time.Sleep(time.Second)
	fmt.Println("exit service")
}

func RunServer(ctx context.Context) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	var cmd *exec.Cmd = nil
	go func() {
		cmd = exec.Command("./webserverdemo1", "-env=local -a=bcd", "-abc=ddd")
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
		fmt.Println("stop cmd command")
	}()
	go func() {

		select {
		case <-ctx.Done():
		}
		cmd.Process.Kill()
		fmt.Println("done")
		wg.Done()
	}()
	wg.Wait()
}
