package main

import (
	"context"
	"fmt"
	"time"
	"os"
	"os/signal"
)

func main() {
	timeOut := 3 * time.Second
	deadline := time.Now().Add(4 * time.Hour)
	timeOutCtx, timeCancel := context.WithTimeout(context.Background(), timeOut)

	// if call timeCancel() will done
	defer timeCancel()


	// cancelFunc会把实现了他子类的所有context 取消掉
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	// 如果把 _ 换成cancelFunc，把上面的 cancelFunc去掉，则只会取消掉他自己。
	cancelCtx2, _ := context.WithCancel(cancelCtx)
	deadlineCtx, _ := context.WithDeadline(cancelCtx, deadline)
	timeOutCtx2, _ := context.WithTimeout(cancelCtx, (time.Hour * 10))

	go contextDemo("[timecount]", timeOutCtx)


	go contextDemo("[cancelContext]", cancelCtx)
	go contextDemo("[cancelContext2]", deadlineCtx)
	go contextDemo("[cancelContext3]", cancelCtx2)
	go contextDemo("[cancelContext4]", timeOutCtx2)
	//
	//go contextDemo("[deadlineContext]", deadlineCtx)


	// Done() will block until it's closed
	// wait for the timeout to expire
	<- timeOutCtx.Done()
	fmt.Println(time.Now())
	// after 5s the timeOutCtx done. the main function will calls the cancelFunc()

	fmt.Println("sleelping")
	time.Sleep(time.Second * 5)
	fmt.Println("stop sleelping begin running")
	//go contextDemo("[cancelContext]", cancelCtx)
	//
	//go contextDemo("[deadlineContext]", deadlineCtx)

	// cancel the deadline context as well as its child -- the cancelCtx

	go func() {
		<- cancelCtx.Done()
		fmt.Println("The cancel context has been cancelled-----" + cancelCtx.Err().Error())

	}()
	go func() {
		fmt.Println("wait for both contexts to be calcelled")
		<- deadlineCtx.Done()
		fmt.Println("the deadline context has been cancelled-----" + cancelCtx.Err().Error())
	}()
	go func() {
		<- timeOutCtx2.Done()
		fmt.Println("The timeOutCtx2 context has been cancelled-----" + timeOutCtx2.Err().Error())

	}()

	fmt.Println("Cancelling the cancel context.....")
	cancelFunc()

	sch := make(chan os.Signal)
	signal.Notify(sch, os.Interrupt)
	<- sch
}


func contextDemo(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():

			fmt.Println(name + "Done.............." + ctx.Err().Error())
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



/*
https://code.tutsplus.com/tutorials/context-based-programming-in-go--cms-29290

 */