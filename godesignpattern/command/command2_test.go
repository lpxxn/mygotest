package command

import "testing"

func TestCommandInvoker(t *testing.T) {
	var com Command = Command{}
	c := ConcreteCommand{"li"}
	*com = c
	com.Executed()

}
