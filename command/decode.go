package command

import (
	"strings"
)

type DecodeCommand struct {
	Meta
}

func (c *DecodeCommand) Run(args []string) int {
	// Write your code here
	return 0
}

func (c *DecodeCommand) Synopsis() string {
	return "Decode Git.io URL"
}

func (c *DecodeCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
