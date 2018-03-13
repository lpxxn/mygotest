package main

import "fmt"

func m1() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
		close(ch1)
	}()

	v, ok := <-ch1
	fmt.Println(ok)
	fmt.Println(v)

	v, ok = <-ch1
	if !ok {
		fmt.Println("read false")
	}
	fmt.Println(v)
}

func m2() {
	ch1 := make(chan int)
	go func() {
		close(ch1)
		// 关闭后不能写入，会报错
		ch1 <- 1
	}()
	fmt.Println(<-ch1)

}
func main() {
	m1()
	m2()
}
