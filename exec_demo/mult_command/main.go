package main

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/mygotest/exec_demo"
)

func main() {
	m := "cd /Users/lipeng/go/src/github.com/mygotest/exec_demo/bad_test; pwd; ls; go list; echo -------"
	cmd := exec.Command("/bin/sh", "-c", m)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			//fmt.Println(scanner.Text())
		}
		fmt.Println("end scan111")
	}()

	//if err := cmd.Start(); err != nil {
	//	panic(err)
	//}
	//
	//if err := cmd.Wait(); err != nil {
	//	panic(err)
	//}
	err = cmd.Run()
	fmt.Printf("err: %#v\n", err)

	cmd = exec.Command("/bin/sh", "-c", "pwd; ls; echo 22222222*****")
	cmd.Dir = "/Users/lipeng/go/src/github.com/mygotest/exec_demo/bad_test"
	cmdReader, err = cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	scanner = bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			//fmt.Println(scanner.Text())
		}
		fmt.Println("end scan2222")
	}()
	err = cmd.Run()
	fmt.Printf("err: %#v\n", err)

	err = exec_demo.RunCmd(m)
	fmt.Printf("err: %#v\n", err)
}
