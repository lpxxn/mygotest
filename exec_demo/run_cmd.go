package exec_demo

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func RunCmd(cmdStr string, cmdDir ...string) error {
	var err error
	fmt.Println("begin run command")
	cmd, stdout, stderr, err := startCmd(cmdStr, cmdDir...)
	if err != nil {
		return err
	}
	defer func() {
		stdout.Close()
		stderr.Close()
	}()
	io.Copy(os.Stdout, stdout)
	io.Copy(os.Stderr, stderr)
	// wait for building
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

func startCmd(cmd string, cmdDir ...string) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	c := exec.Command("/bin/sh", "-c", cmd)
	if len(cmdDir) > 0 && cmdDir[0] != "" {
		c.Dir = cmdDir[0]
	}
	f, err := pty.Start(c)
	return c, f, f, err
}
