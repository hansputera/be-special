package commands

import (
	"strings"

	"github.com/hansputera/be-special/commands/commons"
	"github.com/hansputera/be-special/types"
)

var registeredCommands = []*types.Command{
	commons.PingCommand(),
	commons.StatsCommand(),
}

func FindCommand(cmdOrAlias string) *types.Command {
	cmdOrAlias = strings.ToLower(cmdOrAlias)

	for _, cmd := range registeredCommands {
		if cmd.Name == cmdOrAlias || cmd.ContainAlias(cmdOrAlias) {
			return cmd
		}
	}

	return nil
}
