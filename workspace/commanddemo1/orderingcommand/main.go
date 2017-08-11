package main

import (
	"github.com/urfave/cli"
	"fmt"
	"sort"
	"os"
)
// go run main.go -lang aabcccc a
// go run main.go -l aabcccc a
// go run main.go te add

var la string = ""
func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
			Destination: &la,
			EnvVar: "APP_LAN",
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
				value := c.String("lang")
				fmt.Println(value, la)
				fmt.Println("complete")
				return nil
			},
		},
		{
			Name: "add",
			Aliases: []string{"a"},
			Usage: "add a task to the list",
			Action: Add,
		},
		{
			Name: "temp",
			Aliases:[]string{"te"},
			Usage: "usage subcommands",
			// go run main.go te add
			Subcommands:[]cli.Command{
				{
					Name: "add",
					Usage: "subcommand add",
					Aliases:[]string{"a"},
					Action: func(c *cli.Context) error {
						fmt.Println("subcommand add")
						return nil
					},
				},
				{
					Name: "remove",
					Usage: "subcommand remove",
					Aliases:[]string{"re"},
					Action: func(c *cli.Context) error {
						fmt.Println("subcommand remove")
						return nil
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	//app.Action = func(c *cli.Context) error {
	//	fmt.Println("app Action")
	//	// 这里是可以用的
	//	value := c.String("lang")
	//	fmt.Println(value)
	//	return nil
	//}
	app.Run(os.Args)
}


func Add(c *cli.Context) error {
	// 这里是不可以用的
	// 只能用Destination 的 la
	value := c.String("lang")
	fmt.Println(value, la)
	fmt.Println("add task")
	return nil
}