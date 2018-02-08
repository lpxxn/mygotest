package main

import "fmt"

func main() {
	var f = fmt.Println

	f("Func1 :")
	f(DeferFunc1(10))

	f("Func2 :")
	f(DeferFunc2(10))

	f("Func3 :")
	f(DeferFunc3(10))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()

	return 2

}





