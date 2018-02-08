package main

import "fmt"

func main() {
	f := fmt.Println
	f(a())
	f(b())
	f(c())
}

// 执行顺序为
// 先 return 再 defer

// 作用域为函数体内部
func a() int {
	i := 0
	defer func() {
		i += 1
		fmt.Println("a defer : ", i)
	}()

	return i
}

// 整个函数作用域
func b() (i int) {
	i = 0

	defer func() {
		i += 1
		fmt.Println("b defer : ", i)
	}()
	return i
}

// 和方法b是一样的
// 整个函数作用域
func c() (i int) {
	defer func() {
		i += 1
		fmt.Println("c defer : ", i)
	}()
	// 相当于i = 0
	return 0
}
