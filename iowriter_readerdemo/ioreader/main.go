package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type MyReader struct {
	s string
}

func (r MyReader) Read(b []byte) (int, error) {
	// 如果这样做， io.Copy的时候会报大错
	// strings 的Reader已会把 字符串保存，处理
	//for i := range b {
	//	b[i] = 'A'
	//}
	l := len(b)
	for i := 0; i < l; i++ {
		b[i] = 'A'
	}

	//if r.i >= int64(len(r.s)) {
	//	return 0, io.EOF
	//}
	//r.prevRune = -1
	//n = copy(b, r.s[r.i:])
	//r.i += int64(n)
	//return

	return l, nil
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
	fmt.Println("-----myreader-----")
	//io.Copy(os.Stdout, myreader)

	s3 := strings.NewReader("Abc Pengli")
	r3 := rot13Reader{s3}
	fmt.Println("----rot13Reader---")
	io.Copy(os.Stdout, &r3)
	s3.Seek(io.SeekStart, io.SeekStart)
	fmt.Println("\n----rot13Reader22---")
	io.Copy(os.Stdout, &r3)
	fmt.Println()

	data := []byte("abcdefasdfqweras;dfjkqweroi;lsdjkfasdf")
	beBuf := make([]byte, 4)
	size := uint32(len(data)) + 4

	binary.BigEndian.PutUint32(beBuf, size)
	fmt.Println(beBuf)

	size = 123456
	binary.BigEndian.PutUint32(beBuf, size)
	fmt.Println(beBuf)
	frameType := 2
	binary.BigEndian.PutUint32(beBuf, uint32(frameType))
	fmt.Println(beBuf)

}
