package main

import (
	"fmt"

	"github.com/mygotest/exec_demo"
)

func main() {
	err := exec_demo.RunCmd("go build /Users/lipeng/go/src/github.com/mygotest/exec_demo/bad_test")
	fmt.Printf("err: %#v\n", err)
}
