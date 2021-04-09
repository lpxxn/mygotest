package main_test

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestReq2(t *testing.T) {
	file, _ := os.Open("./progress_req_test.go")
	total := int64(0)
	pr := &ProgressReader{file, func(r int64) {
		total += r
		if r > 0 {
			fmt.Println("progress", r)
		} else {
			fmt.Println("done", r)
		}
	}}
	//io.Copy(ioutil.Discard, pr)
	//io.CopyBuffer(os.Stdout, pr, make([]byte, 100))
	io.Copy(os.Stdout, pr)
}

type ProgressReader struct {
	io.Reader
	Reporter func(r int64)
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.Reader.Read(p)
	pr.Reporter(int64(n))
	return
}
