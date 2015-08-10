package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/spiegel-im-spiegel/gitio-shorten/command"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "ENV_CODE",
		Name:   "code",
		Value:  "",
		Usage:  "Short Code",
	},
}

var Commands = []cli.Command{
	{
		Name:   "encode",
		Usage:  "Shorten GitHub URL",
		Action: command.CmdEncode,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "decode",
		Usage:  "Decode Git.io URL",
		Action: command.CmdDecode,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
