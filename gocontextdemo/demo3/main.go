package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	timeOut := 5 * time.Second
	//deadline := time.Now().Add(4 * time.Hour)
	timeOutCtx, timeCancel := context.WithTimeout(context.Background(), timeOut)
	// if call timeCancel() will done
	defer timeCancel()


	go timeoutFunc("[timecount]", timeOutCtx)


	<- timeOutCtx.Done()
	fmt.Println(time.Now())
}

func timeoutFunc(name string, ctx context.Context) {
	//time.Sleep(2 * time.Second)
	deadline, ok := ctx.Deadline()
	fmt.Println("timeout func")
	if ok {
		fmt.Println(name, "will exire at: ", deadline)
	} else {
		fmt.Println(name, "has no deadline")
	}
	//time.Sleep(time.Second)
}