package main

import (
	"os"
	"runtime/trace"
	"fmt"
	"strconv"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// Your program here
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(strconv.Itoa(i))
		}
	}()

	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}

}
