package whatsapp

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hansputera/be-special/config"
	"github.com/hansputera/be-special/handlers"
	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

func StartWhatsappBot() error {
	container, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?foreign_keys=on", config.BotConfigVar.AuthFile), nil)
	if err != nil {
		return err
	}

	device, err := container.GetFirstDevice()
	if err != nil {
		return err
	}

	client := InitWhatsappClient(device)
	client.AddEventHandler(func(evt interface{}) {
		go handlers.EventHandler(client, evt)
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
	return nil
}
