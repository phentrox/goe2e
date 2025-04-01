package cliFlags

import (
	"context"
	"github.com/phentrox/goseq/cmd"
	"github.com/urfave/cli/v3"
)

func TestDir() cli.Flag {
	return &cli.StringFlag{
		Name:    "dir",
		Aliases: []string{"d"},
		Usage:   "test directory",
		Action: func(ctx context.Context, command *cli.Command, s string) error {
			return cmd.Run(s)
		},
	}

}
