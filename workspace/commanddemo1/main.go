package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)
func main() {
	fmt.Println("hello")

	cli.NewApp().Run(os.Args)
}