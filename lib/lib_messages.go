package lib

import (
	"strings"

	"github.com/hansputera/be-special/config"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

type MessageContext struct {
	Raw    *proto.Message
	Client *whatsmeow.Client
	Info   *types.MessageInfo
}

func (ctx *MessageContext) ExtractPrefix() string {
	subject := strings.ToLower(ctx.Raw.GetConversation())
	for _, prefix := range config.BotConfigVar.Prefixes {
		if strings.HasPrefix(subject, prefix) {
			return prefix
		}
	}

	return ""
}

func (ctx *MessageContext) ExtractRawArgs() []string {
	prefix := ctx.ExtractPrefix()
	if len(prefix) < 1 {
		return []string{}
	}

	text := ctx.Raw.GetConversation()
	return strings.Split(strings.Replace(text, prefix, "", 1), " ")
}

func (ctx *MessageContext) ExtractArgs() []string {
	return ctx.ExtractRawArgs()[1:]
}

func (ctx *MessageContext) GetCommandName() string {
	return ctx.ExtractRawArgs()[0]
}

func ToStringPointer(str string) *string {
	return &str
}
