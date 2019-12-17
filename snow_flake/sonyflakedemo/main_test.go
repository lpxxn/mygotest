package main

import (
	"testing"
	"time"
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

// 单个每分钟生产25601
func TestSonyFlaked2(t *testing.T) {
	var i int
	start := time.Now()
	for ; time.Since(start) < time.Second; {
		sf.NextID()
		i++
	}
	t.Log(i)
}

// 两个同时生产 51201
func TestSonyFlaked3(t *testing.T) {
	consumer := make(chan uint64)
	go func() {
		for {
			id, err := sf2.NextID()
			if err != nil {
				t.Fatal(err)
				return
			}
			consumer <- id
		}
	}()
	go func() {
		for {
			id, err := sf.NextID()
			if err != nil {
				t.Fatal(err)
				return
			}
			consumer <- id
		}
	}()
	count := 0
	start := time.Now()
	for ; time.Since(start) < time.Second; {
		<-consumer
		count++
	}
	t.Log(count)

}
