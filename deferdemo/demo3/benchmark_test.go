package main

import "testing"

func BenchmarkDefer1(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(&t)
	}
}


func BenchmarkDefer2(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doNoDefer(&t)
	}
}


func BenchmarkDefer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoReceiver()
	}
}

/*
go test -v -bench BenchmarkDefer1 -benchmem
go test -test.bench=".*"
 */