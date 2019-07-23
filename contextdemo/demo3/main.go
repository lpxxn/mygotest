package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	tAfter := time.NewTimer(time.Second)
	defer func() {
		cancel()
		tAfter.Stop()
	}()

	go func(ctx context.Context, t *time.Timer, c context.CancelFunc) {
		select {
		case <-ctx.Done():
			println("internal func done")
		case <-t.C:
			c()
			println("process end")
		}
	}(ctx, tAfter, cancel)

	select {
	case <-ctx.Done():
		println("done")
		var err error
		if err = ctx.Err(); err == nil {
			return
		}
		switch err {
		case context.Canceled:
			println("canceled")
		case context.DeadlineExceeded:
			println("deadline exceeded")
		default:
			println(err)
		}

	}
}
