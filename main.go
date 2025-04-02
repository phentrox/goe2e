package main

import (
	"context"
	"github.com/phentrox/goe2e/cmd"
	"github.com/phentrox/goe2e/internal/cli/cliFlags"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	goe2eCommand := cli.Command{
		Name:      "goe2e",
		Usage:     "Go Sequential End-To-End Test Runner",
		UsageText: "goe2e [global options] [command [command options]]",
		Version:   "0.1",
		Commands:  []*cli.Command{},
		Flags: []cli.Flag{
			cliFlags.TestDir(),
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			// only goseq without flags
			// command.Args() does not include flags!
			if len(os.Args) == 1 {
				return cmd.Run("testsE2E")
			}
			return nil // needed so that there is no additional output (main command action always runs!)
		},
	}

	err := goe2eCommand.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
