package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	macro := CommandInvoker{}
	macro.Append(&DrawCommand{PositionInfo: &Position{1, 1}})
	macro.Append(&DrawCommand{PositionInfo: &Position{34, 3}})

	fmt.Println(macro.Execute())

}

func TestCommand2(t *testing.T) {
	macro := CommandInvoker{}
	macro.Append(&DrawCommand{PositionInfo: &Position{41, 31}})
	macro.Append(&DrawCommand{PositionInfo: &Position{324, 3}})

	fmt.Println(macro.Execute())

}
