package main

import (
	"time"
	"fmt"
	"context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go func() {
		dl, ok := ctx.Deadline()
		if ok {
			fmt.Println(dl)
			if time.Until(dl) <= 0  {
				fmt.Println("already exceeded")
				return
			}
		}

		for {
			select {
			case <- ctx.Done():
				fmt.Println("done ctx time.second")
				return

			default:

			}
		}
	}()
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	valueC1 := context.WithValue(ctx, "a", "b")
	valueC2 := context.WithValue(valueC1, "b", "d")
	fmt.Println(valueC2.Value("da"))
	go func() {
		_, err := DoSomeThing(valueC2)
		if err != nil {
			fmt.Println(err)
		}
	}()

	ctx = context.Background()
	_, err := DoSomeThing(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

func DoSomeThing(ctx context.Context) (bool, error){

	workFinish := make(chan struct{})
	go func(ch chan struct{}, ctx context.Context) {
		time.Sleep(time.Second * 5)
		dl, ok := ctx.Deadline()
		if ok {
			fmt.Println(dl)
			if time.Until(dl) <= 0  {
				fmt.Println("DoSomeThing already exceeded")
				return
			}
		}
		fmt.Println("DoSomeThing end.....")
		workFinish <- struct{}{}
	}(workFinish, ctx)

	select {
	case <- workFinish:
		return true, nil
	case <- ctx.Done():
		fmt.Println("DoSomeThing timeout")
		err := ctx.Err()
		return false, err
	}
}
