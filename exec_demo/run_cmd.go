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

func RunSSHCmd(remoteMachine string, cmdStr string) error {
	var err error
	fmt.Println("begin run command")
	cmd, stdout, stderr, err := startSSHCmd(remoteMachine, cmdStr)
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

func startSSHCmd(remoteMachine string, cmd string) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	c := exec.Command("ssh", remoteMachine, cmd)
	f, err := pty.Start(c)
	return c, f, f, err
}

// https://stackoverflow.com/questions/37679939/how-do-i-execute-a-command-on-a-remote-machine-in-a-golang-cli
/*
package main

import (
    "bytes"
    "log"
    "os/exec"
)

func main() {
    cmd := exec.Command("ssh", "remote-machine", "bash-command")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}
To jump over machines use the ProxyCommand directive in a ssh config file.
.ssh/config
Host remote_machine_name
  ProxyCommand ssh -q bastion nc remote_machine_ip 22

*/
