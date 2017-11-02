
package main

import (
	"fmt"
	//"sync/atomic"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	s1 := Stduent{}
	fmt.Println(s1.Speak("aaa"))
	getType(s1)

//	var peo People = Stduent{}
//	think := "bitch"
//	fmt.Println(peo.Speak(think))
}

func getType(in interface{}) {
	switch in.(type) {
	case People :
		fmt.Println("people")
	default:
		fmt.Printf("%T", in)
	}
}
