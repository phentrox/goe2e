package main

import (
	"context"
	"github.com/phentrox/goseq/cmd"
	"github.com/phentrox/goseq/internal/cli/cliFlags"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	goseqCommand := cli.Command{
		Name:      "goseq",
		Usage:     "Go Sequential E2E Test Runner",
		UsageText: "goseq [global options] [command [command options]]",
		Version:   "0.1",
		Commands:  []*cli.Command{},
		Flags: []cli.Flag{
			cliFlags.TestDir(),
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			// only goseq without flags
			// command.Args() does not include flags!
			if len(os.Args) == 1 {
				return cmd.Run("testingE2E")
			}
			return nil // needed so that there is no additional output (main command action always runs!)
		},
	}

	err := goseqCommand.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
