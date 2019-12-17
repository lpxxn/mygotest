package main

import (
	mapset "github.com/deckarep/golang-set"
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

// 单个每秒生产25601
func TestSonyFlaked2(t *testing.T) {
	var i int
	start := time.Now()
	for ; time.Since(start) < time.Second; {
		sf.NextID()
		i++
	}
	t.Log(i)
}

func TestSonyFlaked2S(t *testing.T) {
	var i int
	for end := time.Now().Add(time.Second); ; {
		if time.Now().After(end) {
			break
		}
		sf.NextID()
		i++
	}
	t.Log(i)
}

// 两个同时生产每秒 51201
func TestSonyFlaked3(t *testing.T) {
	// 有没有buffer是一样的结果
	//consumer := make(chan uint64, 10)
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
	for start := time.Now(); time.Since(start) < time.Second; {
		<-consumer
		count++
	}
	t.Log(count)
}

// 检查是否有重复的
func TestSonyFlaked4(t *testing.T) {
	// 有没有buffer是一样的结果
	//consumer := make(chan uint64, 10)
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
	set := mapset.NewSet()
	count := 0
	start := time.Now()
	for ; time.Since(start) < time.Second; {
		idx := <-consumer
		if set.Contains(idx) {
			t.Fatal(" duplicate idx")
		} else {
			set.Add(idx)
		}
		count++
	}
	t.Log(count)
}
