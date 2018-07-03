package main

import (
	"unsafe"
	"fmt"
)

func main() {
	c1 := make(chan []byte, 1000)
	sizeC1 := cap(c1) * int(unsafe.Sizeof(c1))
	fmt.Println(sizeC1)

	b1 := make([]byte, 1000)
	size := cap(b1) * int(unsafe.Sizeof(b1))
	fmt.Println(size)

	iv1 := new(int32)
	sizeIv1 := int(unsafe.Sizeof(*iv1))
	fmt.Println(sizeIv1)

	iv2 := new(int64)
	sizeIv2 := int(unsafe.Sizeof(*iv2))
	fmt.Println(sizeIv2)


	tsa1 := &TestSizeA{Age: 18, Name: "lipeng", Tag: []string{"a", "b"}}
	fmt.Print(tsa1.SizeOTestSizeA())
}


type TestSizeA struct {
	Age int64
	Name string
	Tag []string
}

func (s *TestSizeA) SizeOTestSizeA() int {
	size := int(unsafe.Sizeof(*s))
	size += len(s.Tag)
	return size
}