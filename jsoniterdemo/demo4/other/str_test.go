package main

import "testing"
// go test str_test.go -bench=. -benchmem
func BenchmarkNoPreallocate(b *testing.B) {
	const (
		l = 0
		c = 8 * 1024
	)
	for i := 0; i < b.N; i++ {
		// Don't preallocate our initial slice
		init := []int64{}
		for j := 0; j < c; j++ {
			init = append(init, 42)
		}
	}
}

func BenchmarkPreallocateConst(b *testing.B) {
	const (
		l = 0
		c = 8 * 1024
	)
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, l, c)
		for j := 0; j < cap(init); j++ {
			init = append(init, 42)
		}
	}
}

func BenchmarkPreallocateVar(b *testing.B) {
	var (
		l = 0
		c = 8 * 1024
	)
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, l, c)
		for j := 0; j < cap(init); j++ {
			init = append(init, 42)
		}
	}
}