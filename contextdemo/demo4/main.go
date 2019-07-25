package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "li")
	wg := &sync.WaitGroup{}
	done(ctx, wg)
	fmt.Printf("%#v\n", ctx.Value("name"))
	ctx, cancel := context.WithCancel(ctx)
	done(ctx, wg)
	cancel()

	wg.Wait()
}

func done(ctx context.Context, wg *sync.WaitGroup) {
	done := ctx.Done()
	if done != nil {
		wg.Add(1)
		go func() {
			select {
			case <-done:
				fmt.Println(ctx.Err())
				fmt.Println("done")
			}
			wg.Done()
		}()
	} else {
		fmt.Println("done() chan is empty")
	}
}
