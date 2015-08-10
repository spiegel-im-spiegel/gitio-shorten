package command

import (
	"strings"
)

type EncodeCommand struct {
	Meta
}

func (c *EncodeCommand) Run(args []string) int {
	// Write your code here
	return 0
}

func (c *EncodeCommand) Synopsis() string {
	return "Shorten GitHub URL"
}

func (c *EncodeCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
