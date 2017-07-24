package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello", r)
}

type MyString string

func (s MyString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type MyStruct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s *MyStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", s.Greeting, s.Punct, s.Who)
}

func main() {
	var h Hello
	http.Handle("/string", MyString("I'm a frayed knot."))
	http.Handle("/struct", &MyStruct{"Hello", ":", "Gophers!"})
	err := http.ListenAndServe("localhost:10010", h)

	//err := http.ListenAndServe("localhost:10010", h)

	if err != nil {
		log.Fatal(err)
	}
}
