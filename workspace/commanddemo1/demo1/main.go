package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)
var strv string = ""


var Flags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:"lang, l",
		Value:"eng",
		Usage:"language for the app",
		Destination: &strv,
	},
	cli.StringFlag{
		Name:"lang2, l2",
		Value:"eng",
		Usage:"language for the app",
		Destination: &strv,
	},
}

func main() {
	app := cli.NewApp()

	app.Flags = Flags

	app.Name = "crm"
	app.Usage = "the command used to build and push crm application"
	app.Action = Test
	app.Version = "0.1.1"
	app.Run(os.Args)

}

func Test(c *cli.Context) error{

	name := "Test"
	if c.NArg() > 0 {
		name = c.Args().Get(0)
	}
	fmt.Println(strv, "c.String(`lang`)", c.String("lang"))
	//if c.String("lang") == "eng" {
	if strv == "eng" {
		fmt.Println("Hola", name)
	} else {
		fmt.Println("你好", name)
	}

	return nil
}