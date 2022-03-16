package main_test

import (
	"fmt"
)

type M[K Info, V int64 | string | float64] struct {
	k K
	v V
}

func GetInfo[K Info, V int64 | string | float64](m []M[K, V]) {
	for _, k := range m {
		fmt.Printf("name: %s, age: %d, value: %v \n", k.k.GetName(), k.k.GetAge(), k.v)
	}
}

type Info interface {
	GetName() string
	GetAge() int64
}

type Int1 int

func (i Int1) GetName() string {
	return "int1"
}

func (i Int1) GetAge() int64 {
	return int64(i)
}

type Student struct {
	Name string
	Age  int64
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) GetAge() int64 {
	return s.Age
}
