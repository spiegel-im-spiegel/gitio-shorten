package command

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/gitioapi"
)

type EncodeCommand struct {
	Meta

	Name        string
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}

func (c *EncodeCommand) Run(args []string) int {
	var (
		code string
		url  string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(c.Name, flag.ContinueOnError)
	flags.SetOutput(c.ErrorWriter)
	flags.Usage = func() {}
	flags.StringVar(&code, "code", "", "'code' parameter.")
	flags.StringVar(&code, "c", "", "'code' parameter.")
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

	if len(url) == 0 {
		// Get URL from STDIN
		scanner := bufio.NewScanner(c.Reader)
		if scanner.Scan() {
			url = strings.TrimSpace(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			c.Ui.Error(os.ErrInvalid.Error() + ": " + err.Error())
			return ExitCodeError
		}
	}

	if len(url) == 0 {
		c.Ui.Error(os.ErrInvalid.Error() + ": No Argument")
		return ExitCodeError
	}

	// shortening URL
	shortUrl, err := gitioapi.Encode(&gitioapi.Param{Url: url, Code: code})
	if err != nil {
		c.Ui.Error(err.Error() + "\n")
		return ExitCodeError
	}
	c.Ui.Output(shortUrl)

	return ExitCodeOK
}

func (c *EncodeCommand) Synopsis() string {
	return "Shorten GitHub URL"
}

func (c *EncodeCommand) Help() string {
	helpText := `
usage: gitio-shorten encode [-c <code>] [<url>]
`
	return strings.TrimSpace(helpText)
}
