package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	var langDesc string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &langDesc,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				//  go run clitest.go  -lang=adcde c asde dddde
				lan := langDesc
				fmt.Println(lan, "  args : ", c.Args(), " args length ", len(c.Args()))
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						// go run clitest.go  -lang=adcde t add  ccccc
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
			Action: func(c *cli.Context) error{
				// go run clitest.go  -lang=adcde t cccccccc
				fmt.Println("in template command  !  ")
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println(langDesc, " in action")
		return nil
	}

	app.Run(os.Args)
}
