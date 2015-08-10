package main

import (
	"github.com/mitchellh/cli"
	"github.com/spiegel-im-spiegel/gitio-shorten/command"
)

func Commands(meta *command.Meta, ui *cli.BasicUi) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"encode": func() (cli.Command, error) {
			return &command.EncodeCommand{
				Meta:        *meta,
				Name:        Name,
				Reader:      ui.Reader,
				Writer:      ui.Writer,
				ErrorWriter: ui.ErrorWriter,
			}, nil
		},
		"decode": func() (cli.Command, error) {
			return &command.DecodeCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
