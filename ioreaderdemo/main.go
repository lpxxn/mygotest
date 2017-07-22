package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type MyReader struct {
}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

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

	nb := make([]byte, 10)
	myreader := MyReader{}
	fmt.Println(myreader.Read(nb))
	fmt.Println(string(nb[:]))
}
