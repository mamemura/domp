package main

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name: "query,q",
		Value: "html",
		Usage: "query. useful html/text/attr@xxx, e.g. \"html|attr@id\" , default \"html\"",
	},
	cli.StringFlag{
		Name: "output,o",
		Value: "csv",
		Usage: "output. useful text/csv, default text (feature json) ",
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
