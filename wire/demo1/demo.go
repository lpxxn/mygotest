package main

import (
	"errors"
	"fmt"
	"time"
)

type Message string

type Greeter struct {
	Grumpy  bool
	Message Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage(phrase string) Message {
	return Message(phrase)
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: evnet greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewMessage2(phrase string, b int) Message {
	fmt.Println(b)
	return Message(phrase)
}

func NewEvent2(a string, b int, g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: evnet greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

type Event3Param struct {
	A string
	B float32
	C string
	D int32
}

func NewEvent3(a string, g Greeter, b Event3Param) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: evnet greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}
