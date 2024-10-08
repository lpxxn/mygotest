package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string "User name"
	Age  int    "User Age"
	T    string `w:"100" n:"name"`
}

func main() {

	type Param struct {
		Name *string `json:"name"`
		Age  *int    `json:"age"`
	}

	p := &Param{}
	if err := json.Unmarshal([]byte(`{}`), p); err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", p)
	p = &Param{}
	if err := json.Unmarshal([]byte(`{"name": "abc"}`), p); err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", p)

	p = &Param{}
	if err := json.Unmarshal([]byte(`{"name": "abc", "age": 123}`), p); err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", p)

	user := &User{Name: "Li", Age: 10}

	s := reflect.TypeOf(user).Elem()
	for v, n := 0, s.NumField(); v < n; v++ {
		fmt.Println(s.Field(v).Tag)
	}

	user2 := User{Name: "peng", Age: 100}
	s2 := reflect.TypeOf(user2)
	f2 := s2.Field(0)
	fmt.Println(f2.Tag)
	f3 := s2.Field(2)
	fmt.Println(f3.Tag.Get("weight"), "  n ", f3.Tag.Get("n"))

	m1 := map[User]string{}
	m1[User{Name: "li"}] = "lili"
	m1[User{Name: "li"}] = "peng"
	fmt.Printf("%v\n", m1)
	fmt.Printf("%+v\n", m1)
	fmt.Printf("%#v\n", m1)

	m2 := map[*User]string{}
	m2[&User{Name: "li"}] = "lili"
	m2[&User{Name: "li"}] = "peng"
	fmt.Printf("%+v", m2)
	type sliceStr []string
	sV1 := sliceStr{"a", "b"}
	var sAny any = sV1
	fmt.Println(reflect.TypeOf(sAny))
}
