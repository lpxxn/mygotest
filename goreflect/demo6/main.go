package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

func main() {
	a := struct {}{}
	size1 := unsafe.Sizeof(a)
	fmt.Println("unsafe Sizeof a: ", size1)

	r := reflect.ValueOf(a)
	s2 := binary.Size(r)

	fmt.Println("binary.Size a:", s2)

	b := true
	fmt.Println("bool size of :", unsafe.Sizeof(b))

	i := int64(1)
	fmt.Println("int64 size of :", unsafe.Sizeof(i))

	str2 := ""
	fmt.Println("string size of :", unsafe.Sizeof(str2))


	m1 := make(map[int64]struct{})
	var m_size1, m_size2 runtime.MemStats
	fmt.Println("big map size of :", unsafe.Sizeof(m1))
	runtime.ReadMemStats(&m_size1)

	str1 := "helloworld"
	fmt.Println("string str1 size", unsafe.Sizeof(str1))
	fmt.Println("string str1 size", uintptr(len(str1)) * reflect.TypeOf(str1).Size())

	//bytes1 := []byte(str1)
	bytes1 := make([]byte, len(str1), len(str1))
	for idx, v := range str1 {
		bytes1[idx] = byte(v)
	}
	fmt.Println("bytes1 size", unsafe.Sizeof(bytes1))

	// right
	bytes1Size := uintptr(len(bytes1)) * reflect.TypeOf(bytes1).Elem().Size()
	fmt.Println("bytes1Size size", bytes1Size)
	// right
	fmt.Println("binary.Size of bytes1: ", binary.Size(bytes1))
	// 0.00976563
	fmt.Println("kb", float64(binary.Size(bytes1))/KILOBYTE)

	// suppose  0.2 second
	fmt.Println("mb/s", float64(binary.Size(bytes1))/MEGABYTE/0.2)

	fmt.Println("binary.Size of str1: ", binary.Size(str1))

	for ; i < 5000000; i++ {
		m1[i] = struct{}{}
	}
	runtime.ReadMemStats(&m_size2)
	fmt.Println("big map size of :", unsafe.Sizeof(m1))
	memUsage(&m_size1, &m_size2)

	intTemp2 := int64(1)
	fmt.Println("Size of []int64:", unsafe.Sizeof(intTemp2))
	fmt.Println("Size of [5000000]int64:", unsafe.Sizeof([5000000]int64{}))
	fmt.Println("Size of [5000000]struct:", unsafe.Sizeof([5000000]struct{}{}))
}

var p = fmt.Println

func memUsage(m1, m2 *runtime.MemStats) {
	p("Alloc:", m2.Alloc-m1.Alloc,
		"TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc,
		"HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}