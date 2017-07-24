package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = p[i] + 1
	}
	return
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

	s3 := strings.NewReader("Abc Pengli")
	r3 := rot13Reader{s3}
	io.Copy(os.Stdout, &r3)
}
