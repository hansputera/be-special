package cmd

import (
	"github.com/hansputera/be-special/config"
	"github.com/urfave/cli/v2"
)

func beforeActHandler(ctx *cli.Context) error {
	err := config.InitConfig(ctx.String("config"))
	if err != nil {
		return err
	}

	return err
}
