package command

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/gitioapi"
)

type DecodeCommand struct {
	Meta

	Name        string
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}

func (c *DecodeCommand) Run(args []string) int {
	var url string

	// Define option flag parse
	flags := flag.NewFlagSet(c.Name, flag.ContinueOnError)
	flags.SetOutput(c.ErrorWriter)
	flags.Usage = func() {}
	// Parse commandline flag
	if err := flags.Parse(args); err != nil {
		return ExitCodeError
	}
	// Parse argument
	switch flags.NArg() {
	case 0:
		url = ""
	case 1:
		url = flags.Arg(0)
	default:
		c.Ui.Error(os.ErrInvalid.Error() + ": " + flags.Arg(1))
		return ExitCodeError
	}

	// shortening URL
	decodeUrl, err := gitioapi.Decode(&gitioapi.Param{Url: url})
	if err != nil {
		c.Ui.Error(err.Error() + "\n")
		return ExitCodeError
	}
	c.Ui.Output(decodeUrl)

	return ExitCodeOK
}

func (c *DecodeCommand) Synopsis() string {
	return "Decode Git.io URL"
}

func (c *DecodeCommand) Help() string {
	helpText := `
usage: gitio-shorten decode <url>
`
	return strings.TrimSpace(helpText)
}
