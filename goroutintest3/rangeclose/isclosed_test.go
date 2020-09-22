package main

import (
	"fmt"
	"testing"
	"time"
)

func Benchmark_IsClosed(b *testing.B) {
	ch1 := make(chan string, 2)
	ch1 <- "a"
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		isChanClosed(ch1)

	}
}

type A struct {
	Name string
}

func TestR(t *testing.T) {
	a := []*A{{Name: "a"}, {"b"}, {"c"}}
	for _, item := range a {
		go ft(item)
	}
	time.Sleep(time.Second * 2)
}
func ft(a *A) {
	fmt.Println(*a)
}
