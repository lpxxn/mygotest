package main

import (
	"reflect"
	"fmt"
)

type User struct {
	Name string "User name"
	Age int "User Age"
	T string `w:"100" n:"name"`
}

func main() {
	user := &User{Name:"Li", Age: 10}

	s := reflect.TypeOf(user).Elem()
	for v, n := 0, s.NumField(); v < n; v++ {
		fmt.Println(s.Field(v).Tag)
	}

	user2 := User{Name:"peng", Age:100}
	s2:= reflect.TypeOf(user2)
	f2 := s2.Field(0)
	fmt.Println(f2.Tag)
	f3 := s2.Field(2)
	fmt.Println(f3.Tag.Get("weight"), "  n ", f3.Tag.Get("n"))

}
