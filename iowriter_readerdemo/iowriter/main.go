package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func (p *Person) Write(w io.Writer) (n int, err error) {
	b, _ := json.Marshal(p)
	return w.Write(b)
}

func main() {
	me := Person{
		Id:   1,
		Name: "H",
		Age:  10,
	}
	var b bytes.Buffer
	me.Write(&b)
	fmt.Println(b.String())

	bWriter := bufio.NewWriterSize(os.Stdout, 10)
	bWriter.WriteString("12345678910\n") //因为size是 10，大于size就Flush(); 会先输出1234567891

	bWriter.WriteString("abc\n")
	bWriter.WriteString("def\n")
	bWriter.WriteString("higklmn\n")
	bWriter.Flush()
	/* 如果没有调用Flush() 方法会只输出下面这些
	12345678910
	abc
	def
	*/
}
