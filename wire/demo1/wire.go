//+build wireinject

package main

import (
	"fmt"
	"github.com/google/wire"
)

var S string

func TestMsg(s string) string {
	S = s
	fmt.Println(S)
	return s
}

func InitializeEvent(phrase string) (Event, error) {
	//wire.Build(NewEvent, NewGreeter, NewMessage)
	//wire.Build(NewMessage, NewEvent, NewGreeter, TestMsg)
	wire.Build(NewGreeter, NewMessage, NewEvent)
	return Event{}, nil
}
