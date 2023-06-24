package messages

import (
	"context"

	"github.com/hansputera/be-special/commands"
	"github.com/hansputera/be-special/lib"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

func MessagesHandler(client *whatsmeow.Client, m *events.Message) {
	messageContext := &lib.MessageContext{
		Raw:    m.Message,
		Client: client,
		Info:   &m.Info,
	}

	cmd := commands.FindCommand(messageContext.GetCommandName())
	if cmd != nil {
		err := cmd.Action(messageContext)
		if err != nil {
			client.SendMessage(context.Background(), m.Info.Chat, &proto.Message{
				ExtendedTextMessage: &proto.ExtendedTextMessage{
					ContextInfo: &proto.ContextInfo{
						QuotedMessage: m.Message,
						StanzaId:      &m.Info.ID,
					},
					Text: lib.ToStringPointer(err.Error()),
				},
			})
		}
	}
}
