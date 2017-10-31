package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			// set alternate (or short) names for flags by providing a comma-delimited list for the Name e.g.
			Name: "config, c",
			// Note that oly the first placeholder is used. Subsequent back-quoted words will be left as-is
			Usage: "Load configuration from `File aa` `bb`",
		},
	}
	app.Run(os.Args)
}
