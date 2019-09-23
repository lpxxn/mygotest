package main

import (
	"fmt"

	"github.com/mygotest/tree"
)

func walkImpl(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch, quit)

	select {
	case ch <- t.Value:
		fmt.Println("t: ", t.Value, "ch :", ch)
		// Value successfully sent.
	case <-quit:
		return
	}
	walkImpl(t.Right, ch, quit)
}

func walk(t *tree.Tree, ch, quit chan int) {
	walkImpl(t, ch, quit)
	close(ch)
}

func same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go walk(t1, w1, quit)
	go walk(t2, w2, quit)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		fmt.Println("reslut : ok1, ok2", v1, " v2 ", v2)
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	t1, t2 := tree.New(1), tree.New(1)
	fmt.Println("t1 :", t1, "  ", " t2 :", t2)
	if same(t1, t2) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	t1 = tree.New(1)
	t2 = tree.New(2)
	fmt.Println("t1 :", t1, "  ", " t2 :", t2)
	if !same(t1, t2) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
