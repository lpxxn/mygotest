package main

import "fmt"

func main() {
	ch1 := make(chan *int, 1)
	v1 := 1
	ch1 <- &v1

	var vint *int
	select {
	case vint = <- ch1:
		fmt.Println("run <-ch1")
	default:
		fmt.Println("run default")
		v2 := 2323
		vint = &v2
	}
	fmt.Println(vint, *vint)
	select {
	case vint = <- ch1:
		fmt.Println("run <-ch1")
	default:
		fmt.Println("run default")
		v2 := 2323
		vint = &v2
	}
	fmt.Println(vint, *vint)
}
