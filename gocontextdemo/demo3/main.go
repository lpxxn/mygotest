package main

import (
	"fmt"
	"context"
	"time"
	"os"
	"os/signal"
)

func main() {
	timeOut := 5 * time.Second
	//deadline := time.Now().Add(4 * time.Hour)
	timeOutCtx, timeCancel := context.WithTimeout(context.Background(), timeOut)
	// if call timeCancel() will done
	defer timeCancel()

	//timeCancel()

	go timeoutFunc("[timecount]", timeOutCtx)

	//time.AfterFunc(time.Second + 2, func() {
	//	timeCancel()
	//})


	ch, ok := <- timeOutCtx.Done()
	fmt.Println(ch, ok)
	fmt.Println(time.Now())
	sch := make(chan os.Signal)
	signal.Notify(sch, os.Interrupt)

	fmt.Println("-====")
	ch1 := make(chan struct{})
	close(ch1)
	v1, ok1 := <-ch1
	v2, ok2 := <-ch1
	fmt.Println(v1, v2, " ok ", ok1 ,ok2 )
	<-sch
}

func timeoutFunc(name string, ctx context.Context) {
	for {
		select {
		case ch, ok := <-ctx.Done():
			fmt.Println(name + "Done.............." + ctx.Err().Error(), "  ch ", ch, "  ok :", ok)
			return

		default:
			deadline, ok := ctx.Deadline()
			if ok {
				fmt.Println(name, "will exire at: ", deadline)
			} else {
				fmt.Println(name, "has no deadline")
			}
		}
		time.Sleep(time.Second)
	}
}