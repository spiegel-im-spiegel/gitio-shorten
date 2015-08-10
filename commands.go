package main

import (
	"github.com/mitchellh/cli"
	"github.com/spiegel-im-spiegel/gitio-shorten/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"encode": func() (cli.Command, error) {
			return &command.EncodeCommand{
				Meta: *meta,
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
