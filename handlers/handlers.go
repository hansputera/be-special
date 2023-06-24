package handlers

import (
	"github.com/hansputera/be-special/handlers/messages"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(client *whatsmeow.Client, evt interface{}) {
	switch event := evt.(type) {
	case *events.Message:
		go messages.MessagesHandler(client, event)
	}
}