package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/creack/pty"
	"github.com/fsnotify/fsnotify"
)

func main() {
	var exitCh = make(chan os.Signal)
	ctx, cancel := context.WithCancel(context.Background())

	signal.Notify(exitCh, os.Interrupt)

	//go RunServer(ctx)
	//go RunServerContext(ctx)
	go RunServerContextWatch(ctx)

	fmt.Println("run service")
	<-exitCh
	cancel()
	time.Sleep(time.Second)
	fmt.Println("exit service")
}

func RunServer(ctx context.Context) {
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
	<-ctx.Done()
	cmd.Process.Kill()
	fmt.Println("done")
}

func RunServerContext(ctx context.Context) {
	var cmd *exec.Cmd = nil
	go func() {
		cmd = exec.CommandContext(ctx, "./webserverdemo1", "-env=local -a=bcd", "-abc=ddd")
		cmd.Env = append(cmd.Env, "ENV=dev")
		fmt.Println(cmd.String())
		fmt.Println(cmd.Env)
		f, err := pty.Start(cmd)
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, f)
		fmt.Println("stop cmd command")
	}()
	<-ctx.Done()
	fmt.Println("done")
}

func RunServerContextWatch(ctx context.Context) {

	watchFile := "./webserverdemo1"

	runAndWatch(watchFile)

	<-ctx.Done()

	fmt.Println("done")
}

func runAndWatch(watchFile string) {
	var cmd *exec.Cmd = nil

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	cmdCtx, cancel := context.WithCancel(context.Background())
	go func() {
		cmd = exec.CommandContext(cmdCtx, "./webserverdemo1", "-env=local -a=bcd", "-abc=ddd")
		cmd.Env = append(cmd.Env, "ENV=dev")
		fmt.Println(cmd.String())
		fmt.Println(cmd.Env)
		f, err := pty.Start(cmd)
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, f)
		fmt.Println("stop cmd command")
	}()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					watcher.Remove(watchFile)
					watcher.Close()
					cancel()
					fmt.Println("file removed restart service")
					time.Sleep(time.Second * 2)
					runAndWatch(watchFile)
				}
			case err := <-watcher.Errors:
				if err != nil {
					fmt.Println("error:", err)
					panic(err)
				}
			}
		}
	}()

	err = watcher.Add(watchFile)
	if err != nil {
		cancel()
		panic(err)
	}
}

/*
https://github.com/fsnotify/fsnotify/blob/master/example_test.go
https://github.com/howeyc/fsnotify
*/
