package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestEncodeCommand_implement(t *testing.T) {
	var _ cli.Command = &EncodeCommand{}
}
