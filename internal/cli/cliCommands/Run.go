package cliCommands

import (
	"context"
	"github.com/phentrox/goseq/cmd"
	"github.com/urfave/cli/v3"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "runs tests",
		Action: func(ctx context.Context, command *cli.Command) error {
			return cmd.Run("testing/testingE2E")
		},
	}
}
