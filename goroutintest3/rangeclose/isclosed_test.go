package main

import "testing"

func Benchmark_IsClosed(b *testing.B) {
	ch1 := make(chan string, 2)
	ch1 <- "a"
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		isChanClosed(ch1)

	}
}