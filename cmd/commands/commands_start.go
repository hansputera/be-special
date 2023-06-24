package commands

import (
	"github.com/hansputera/be-special/whatsapp"
	"github.com/urfave/cli/v2"
)

func StartCommand(ctx *cli.Context) error {
	return whatsapp.StartWhatsappBot()
}