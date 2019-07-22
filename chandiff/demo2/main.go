package main

import (
	"fmt"
	"sync"
)

type Ct1 struct {
	A chan int
}

type Ct2 struct {
	B chan int
}

func main() {
	newCh := make(chan int)

	c1 := Ct1{A: newCh}

	c2 := Ct2{B: c1.A}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go getValue(c1.A, 1, &wg)

	go func() {
		c2.B <- 1
		c1.A <- 2

		close(c2.B)
		// error
		// c1.A <- 3
		getValue(c1.A, 2, &wg)
	}()

	for v := range newCh {
		fmt.Println("main func get value: ", v)
	}
	if v, ok := <-newCh; ok {
		fmt.Println("newCh value", v)
	}
	if _, ok := <-c1.A; !ok {
		fmt.Println("C1: A is closed")
	}
	fmt.Println("wait--------")
	wg.Wait()
}

func getValue(ch1 <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	v, ok := <-ch1
	fmt.Println("id: ", id)
	if !ok {
		fmt.Println("not ok already closed, v:", v)
		return
	}
	fmt.Println("getValue value: ", v)

}

/*

id:  1
main func get value:  1
C1: A is closed
wait--------
getValue value:  2
id:  2
not ok already closed, v: 0


id:  2
main func get value:  2
not ok already closed, v: 0
C1: A is closed
wait--------
id:  1
getValue value:  1

*/
