package command

import "fmt"

// command interface
type Command interface {
	Executed()
}

type ConcreteCommand struct {
	Name string
}

func (concrete *ConcreteCommand) Executed() {
	fmt.Println(concrete.Name, "Hello")
}

// Invoker
type Invoker struct {
	Comands []Command
}

func (invoker *Invoker) SetCommand(command *Command) {
	invoker.Comands = append(invoker.Comands, command)
}

func (invoker *Invoker) Execute() {
	for _, item := range invoker.Comands {
		item.Executed()
	}
}
