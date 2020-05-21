package main

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/mygotest/exec_demo"
)

func main() {
	// 最后的 gobuild 如果失败 下面都会报错.但 cmd.Start()不会
	m := "cd /Users/lipeng/go/src/github.com/mygotest/exec_demo/bad_test; pwd; ls; go list; echo -------; go build"
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
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd Run error: %#v", err)
		//panic(err)
	}
	//if err := cmd.Start(); err != nil {
	//	panic(err)
	//}
	//
	//if err := cmd.Wait(); err != nil {
	//	panic(err)
	//}

	err = exec_demo.RunCmd(m)
	fmt.Printf("exec_demo.RunCmd err: %#v\n", err)
}
