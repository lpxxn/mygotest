package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	timeOut := 5 * time.Second
	deadline := time.Now().Add(4 * time.Hour)
	timeOutCtx, timeCancel := context.WithTimeout(context.Background(), timeOut)
	// if call timeCancel() will done
	defer timeCancel()


	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	deadlineCtx, _ := context.WithDeadline(cancelCtx, deadline)


	go contextDemo(context.WithValue(
		timeOutCtx, "name", "[timeoutContext]"))
	go contextDemo(context.WithValue(
		cancelCtx, "name", "[cancelContext]"))
	go contextDemo(context.WithValue(
		deadlineCtx, "name", "[deadlineContext]"))


	// Done() will block until it's closed
	// wait for the timeout to expire
	<- timeOutCtx.Done()
	// after 5s the timeOutCtx done. the main function will calls the cancelFunc()

	//go contextDemo("[cancelContext]", cancelCtx)
	//
	//go contextDemo("[deadlineContext]", deadlineCtx)

	// cancel the deadline context as well as its child -- the cancelCtx
	fmt.Println("Cancelling the cancel context.....")
	cancelFunc()

	<- cancelCtx.Done()
	fmt.Println("The cancel context has been cancelled")

	fmt.Println("wait for both contexts to be calcelled")
	<- deadlineCtx.Done()
	fmt.Println("the deadline context has been cancelled....")

}


func contextDemo(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	name := fmt.Sprint(ctx.Value("name"))

	for {
		if ok {
			fmt.Println(name, " will expire ad :", deadline)
		} else {
			fmt.Println(name, " has no deadline")
		}
		time.Sleep(time.Second)
	}
}