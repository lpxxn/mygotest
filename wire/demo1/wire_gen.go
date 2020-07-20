// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"fmt"
)

// Injectors from abcde.go:

func InitializeEvent(phrase string) (Event, error) {
	message := NewMessage(phrase)
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func Test(phrase string) (Event, error) {
	message := NewMessage(phrase)
	greeter := NewGreeter(message)
	event, err := NewEvent(greeter)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func InitEvent2(a string, b int) (Event, error) {
	message := NewMessage2(a, b)
	greeter := NewGreeter(message)
	event, err := NewEvent2(a, b, greeter)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

// abcde.go:

var S string

func TestMsg(s string) string {
	S = s
	fmt.Println(S)
	return s
}
