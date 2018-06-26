package main

import "context"

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for { //for循环是goroutine 函数作为服务的标志
				select {
				case <-ctx.Done():
					close(dst)
					return
				case dst <- n: //这里堵塞有可能会成为函数的超时依赖
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //一定要手动加取消操作
	for n := range gen(ctx) {
		println(n)
		if n == 5 {
			break
		}
		if n == 3 {
			cancel()
		}
	}
}
