package types

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
)

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Action      func(client *whatsmeow.Client, m *proto.Message) error

	OnlyGroup bool
	OnlyPM    bool
}
