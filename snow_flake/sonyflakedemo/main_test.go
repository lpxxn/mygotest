package main

import (
	"testing"
)

func BenchmarkSonyFlaked(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sf.NextID()
	}
}

func BenchmarkSonyFlaked2(b *testing.B) {
	consumer := make(chan uint64)

	go func() {
		for {
			id, err := sf.NextID()
			if err != nil {
				b.Fatal(err)
				return
			}
			consumer <- id
		}
	}()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-consumer
	}
	b.StopTimer()
	b.Log(b.N)
}

func TestSonyFlaked1(t *testing.T) {
	for i := 0; i < 200; i++ {
		sf.NextID()
	}
}
