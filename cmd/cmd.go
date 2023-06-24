package cmd

import (
	"log"
	"os"

	"github.com/hansputera/be-special/cmd/commands"
	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:  "be-special",
		Usage: "I can't describe what the fuck is it",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c", "cfg"},
				Usage:    "File configuration path",
				Required: true,
			},
		},
		Before: beforeActHandler,
		Commands: []*cli.Command{
			{
				Name: "start",
				Usage: "Start the bot",
				Aliases: []string{"run", "ran"},
				Action: commands.StartCommand,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
