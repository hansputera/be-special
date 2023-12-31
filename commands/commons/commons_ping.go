package commons

import (
	"context"

	"github.com/hansputera/be-special/lib"
	"github.com/hansputera/be-special/types"
	"go.mau.fi/whatsmeow/binary/proto"
)

func PingCommand() *types.Command {
	return &types.Command{
		Name:        "ping",
		Description: "Ping/test command to test the bot",
		Aliases:     []string{"test"},
		Action: func(ctx *lib.MessageContext) error {
			ctx.Client.SendMessage(context.Background(), ctx.Info.Chat, &proto.Message{
				ExtendedTextMessage: &proto.ExtendedTextMessage{
					ContextInfo: &proto.ContextInfo{
						StanzaId:      &ctx.Info.ID,
						QuotedMessage: ctx.Raw,
					},
					Text: lib.ToStringPointer("Pong!"),
				},
			})
			return nil
		},
	}
}
