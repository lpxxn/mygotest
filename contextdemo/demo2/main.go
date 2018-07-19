package main

import (
	"time"
	"fmt"
	"context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go func() {
		time.Sleep(time.Second * 10)
		cancel()
	}()
	_, err := DoSomeThing(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

func DoSomeThing(ctx context.Context) (bool, error){

	workFinish := make(chan struct{})
	go func(ch chan struct{}) {
		time.Sleep(time.Second * 10)
		fmt.Print("DoSomeThing end.....")
		workFinish <- struct{}{}
	}(workFinish)

	select {
	case <- workFinish:
		return true, nil
	case <- ctx.Done():
		err := ctx.Err()
		return false, err
	}
}
