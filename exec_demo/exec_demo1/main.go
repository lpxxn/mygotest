package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func main() {
	err := runCmd("GOOS=linux go build /Users/lipeng/go/src/github.com/mygotest/exec_demo/bad_test")
	fmt.Printf("err: %#v\n", err)

}
func runCmd(cmdStr string) error {
	var err error
	fmt.Println("begin run command")
	cmd, stdout, stderr, err := startCmd(cmdStr)
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

func startCmd(cmd string) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	c := exec.Command("/bin/sh", "-c", cmd)
	f, err := pty.Start(c)
	return c, f, f, err
}
