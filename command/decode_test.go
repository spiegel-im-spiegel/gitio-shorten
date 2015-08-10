package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDecodeCommand_implement(t *testing.T) {
	var _ cli.Command = &DecodeCommand{}
}
