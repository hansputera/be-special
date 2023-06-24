package types

import (
	"github.com/hansputera/be-special/lib"
)

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Action      func(ctx *lib.MessageContext) error

	OnlyGroup bool
	OnlyPM    bool
}

func (c *Command) ContainAlias(alias string) bool {
	for _, aliase := range c.Aliases {
		if aliase == alias {
			return true
		}
	}

	return false
}
