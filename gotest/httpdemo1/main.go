package main

import (
	"fmt"
	"net/http"
)

type PersonInfo interface {
	Name() string
	SetName(name string)
}

type Student struct {
	name string
}

func (s *Student) Name() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

func NewStudent(name string) *Student {
	return &Student{name: name}
}

func NameHandler(p PersonInfo) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Hello! My name is %s!", p.Name())
	}
}

func SetName(p PersonInfo) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		name := r.URL.Query().Get("name")
		p.SetName(name)
		fmt.Fprintf(w, "name is update to %s", name)
	}
}

//func SetNameHandler(p PersonInfo) http.HandlerFunc

func main() {
	student := NewStudent("li")
	http.HandleFunc("/name", NameHandler(student))
	http.HandleFunc("/setName", SetName(student))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("service stopping....")
}
