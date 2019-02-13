package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
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

	m1 := make(map[int64]struct{})
	var m_size1, m_size2 runtime.MemStats
	fmt.Println("big map size of :", unsafe.Sizeof(m1))
	runtime.ReadMemStats(&m_size1)

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