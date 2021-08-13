package main

import "testing"
// https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d
func TestSliceStr(t *testing.T) {
	s1 := make([]byte, 0, 2)
	t.Log(&s1)
	t.Log(cap(s1))
	s1 = append(s1, []byte("adfasfwersdfawerqwfadflasjfashkfhsf一二三四五六七")...)
	t.Log(cap(s1))
	t.Log(string(s1))
	var (
		a = 0
		b = 2
	)
	s2 := make([]byte, a, b)
	t.Log(&s2)
}
// https://stackoverflow.com/questions/47192729/interpretting-benchmarks-of-preallocating-a-slice
/*
Go has an optimizing compiler. Constants are evaluated at compile time. Variables are evaluated at runtime.
Constant values can be used to optimize compiler generated code. For example,

The benchmark tool only reports heap allocations. Stack allocations via escape analysis are less costly, possibly free, so are not reported.
The gcflag -m flag appears to show info that shows whether the var is assigned in the stack or heap

All goroutines share a common heap and anything that can’t be stored on the stack will end up there.
When a heap allocation occurs in a function being benchmarked, we’ll see the allocs/ops stat go up by one.
It’s the job of the garbage collector to later free heap variables that are no longer referenced.
*/
// go test str_test.go -bench=. -benchmem
// benchmark slice
func BenchmarkNoPreallocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Don't preallocate our initial slice
		init := []int64{}
		init = append(init, 5)
	}
}

func BenchmarkPreallocateConst(b *testing.B) {
	const (
		l = 0
		c = 1
	)
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, l, c)
		init = append(init, 5)
	}
}

func BenchmarkPreallocateVar(b *testing.B) {
	var (
		l = 0
		c = 1
	)
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, l, c)
		init = append(init, 5)
	}
}

