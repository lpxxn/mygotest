package main

import (
	"fmt"
	"strings"
)

type f1 func(a string) string

func (f f1) Bind(s string) string{
	return f(s + " li")
}

var print = fmt.Println
func main() {

	var f f1 = func(a string) string {
		return a + "hello"
	}
	r := f.Bind("peng")
	print(r)
	s1 := "abcde"
	strings.Compare("a", "b")
}