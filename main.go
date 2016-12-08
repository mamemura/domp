package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "makoto"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Action = CmdDomCutter
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
