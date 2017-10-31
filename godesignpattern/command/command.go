package command

import "strconv"

type Commander interface {
	Execute() string
}

type CommandInvoker struct {
	commands []Commander
}

func (self *CommandInvoker) Execute() string {
	var result string
	for _, command := range self.commands {
		result += command.Execute() + "\n"
	}
	return result
}

func (self *CommandInvoker) Append(command Commander) {
	self.commands = append(self.commands, command)
}

func (self *CommandInvoker) Undo() {
	if len(self.commands) != 0 {
		self.commands = self.commands[:len(self.commands)-1]
	}
}

// receiver
type Position struct {
	X, Y int
}

func (p *Position) Point() string {
	return strconv.Itoa(p.X) + "." + strconv.Itoa(p.Y)
}

// Concrete Command
type DrawCommand struct {
	PositionInfo *Position
}

func (self *DrawCommand) Execute() string {
	return self.PositionInfo.Point()
}
