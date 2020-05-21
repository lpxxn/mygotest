package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "pwd; ls; echo 22222222*****; go build;")
	cmd.Dir = "/Users/lipeng/go/src/github.com/mygotest/exec_demo/bad_test"
	sOut, err := cmd.StdoutPipe()
	sErr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	//TODO: 注意 这个 err 是 nil
	err = cmd.Start()
	fmt.Printf("cmd.Start err: %#v\n", err)
	s, err := ioutil.ReadAll(sOut)
	fmt.Printf("body out : %s ,err: %#v\n", string(s), err)
	s, err = ioutil.ReadAll(sErr)
	fmt.Printf("body err: %s ,err: %#v\n", string(s), err)
	//err = exec_demo.RunCmd(m)
	fmt.Printf("err: %#v\n", err)
}
