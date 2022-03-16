package main_test

import (
	"fmt"
	"testing"
)

func TestGeneric1(t *testing.T) {
	v1 := []M[Info, int64]{
		{k: Int1(1), v: 123},
		{k: Int1(7), v: 345},
	}
	v2 := []M[Info, string]{
		{k: Student{Name: "li", Age: 1}, v: "hello"},
		{k: Student{Name: "zhang", Age: 2}, v: "world"},
	}
	for _, k := range v1 {
		fmt.Printf("name: %s, age: %d, value: %v \n", k.k.GetName(), k.k.GetAge(), k.v)
	}
	GetInfo(v1)
	GetInfo(v2)
	// 用map要把 type Info interface 里的 // comparable 去掉
	//fmt.Println("--------------")
	//m1 := map[Int1]int64{1: 123, 2: 246}
	//GetInfo2(m1)
	//
	//m2 := map[Student]string{Student{Name: "li", Age: 1}: "hello", Student{Name: "zhang", Age: 2}: "world"}
	//GetInfo2(m2)
}

type vInterface interface {
	int64 | string | float64
}

type M[K Info, V vInterface] struct {
	k K
	v V
}

func GetInfo[K Info, V vInterface](m []M[K, V]) {
	for _, k := range m {
		fmt.Printf("name: %s, age: %d, value: %v \n", k.k.GetName(), k.k.GetAge(), k.v)
	}
}

func GetInfo2[K Info, V vInterface](m map[K]V) {
	for k, v := range m {
		fmt.Printf("name: %s, age: %d, value: %v\n", k.GetName(), k.GetAge(), v)
	}
}

type Info interface {
	// comparable // 用map
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
