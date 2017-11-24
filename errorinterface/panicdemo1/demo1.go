package main

import "fmt"

func main() {
	tchan := make(chan int)

	go TesFun(tchan)
	fmt.Println("running")

	var a int = <-tchan
	fmt.Println(a)
}

func TesFun(tchan chan<- int) {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("defer recover err : ", err)
			tchan <- 3
		}
	}()

	fmt.Println("haha")
	//	panic(nil) 这种是不会被 defer里的recover抓住 会 deadlock
	//panic(nil)
	tchan <- 1
}
