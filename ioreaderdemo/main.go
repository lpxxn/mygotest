package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader !")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)

		if err == io.EOF {
			break
		}
		fmt.Printf("n = %v err = %v b = %v \n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
	}
	fmt.Printf("b = %q\n", b)
	r2 := strings.NewReader("Hello, Reader !")
	if data, err := ioutil.ReadAll(r2); err == nil {
		fmt.Println("data", data)
		fmt.Printf("data = %q\n", data)
	} else {
		fmt.Println("err", err)
	}
}
