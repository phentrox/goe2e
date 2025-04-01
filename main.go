package main

import (
	"context"
	"github.com/phentrox/goseq/internal/cli/cliCommands"
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
		Commands: []*cli.Command{
			cliCommands.Run(),
		},
		Flags: []cli.Flag{
			cliFlags.TestDir(),
		},
	}

	err := goseqCommand.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
