package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
var f = fmt.Println
// it is possible to close a non-empty channel but still have the remaining values be received.
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	f(isChanClosed(queue))
	close(queue)


	f(isChanClosed(queue))

	//v , ok := <-queue
	//
	//f(v, " ok: ", ok)

	//close(queue)


	for elem := range queue {
		f(elem)
	}

	for ele := range queue{
		f("abc: ", ele)

	}

	v , ok := <-queue

	f(v, " ok: ", ok)
}


func isChanClosed(ch interface{}) bool {
	if reflect.TypeOf(ch).Kind() != reflect.Chan {
		panic("only channels!")
	}

	// get interface value pointer, from cgo_export
	// typedef struct { void *t; void *v; } GoInterface;
	// then get channel real pointer
	cptr := *(*uintptr)(unsafe.Pointer(
		unsafe.Pointer(uintptr(unsafe.Pointer(&ch)) + unsafe.Sizeof(uint(0))),
	))

	// this function will return true if chan.closed > 0
	// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
	// type hchan struct {
	// qcount   uint           // total data in the queue
	// dataqsiz uint           // size of the circular queue
	// buf      unsafe.Pointer // points to an array of dataqsiz elements
	// elemsize uint16
	// closed   uint32
	// **

	cptr += unsafe.Sizeof(uint(0))*2
	cptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))
	cptr += unsafe.Sizeof(uint16(0))
	return *(*uint32)(unsafe.Pointer(cptr)) > 0
}