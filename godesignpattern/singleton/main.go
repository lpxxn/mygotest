package main

import (
	"sync"
	"fmt"
)

type singleton struct {
	Name string
}

var (
	once sync.Once

	instance *singleton
)

func New() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}

func main() {
	s1 := New()
	s1.Name = "li"

	s2 := New()
	fmt.Println(s2)
}
