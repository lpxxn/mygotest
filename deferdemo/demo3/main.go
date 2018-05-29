package main

import "fmt"

func doNoDefer(t *int) {
	func() {
		*t++
	}()
}
func doDefer(t *int) {
	defer func() {
		*t++
	}()
}


func DoReceiver() {
	defer func(){
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}

