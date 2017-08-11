package main

import (
	"github.com/urfave/cli"
	"fmt"
	"sort"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name: "config, c",
			Usage: "Load configuration form `File`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "complete",
			Aliases: []string{"c"},
			Usage: "complete a task onthe list",
			Action: func(c *cli.Context) error {
				fmt.Println("complete")
				return nil
			},
		},
		{
			Name: "add",
			Aliases: []string{"a"},
			Usage: "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("add task")
				return nil
			},
		},

	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Action = func(c *cli.Context) error {
		fmt.Println("app Action")
		return nil
	}
	app.Run(os.Args)
}
