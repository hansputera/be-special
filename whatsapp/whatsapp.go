package whatsapp

import (
	"context"
	"log"

	"github.com/hansputera/be-special/config"
	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func InitWhatsappClient(device *store.Device) *whatsmeow.Client {
	client := whatsmeow.NewClient(device, waLog.Stdout("WAClient", config.BotConfigVar.LogLevel, true))

	if client.Store.ID == nil {
		qrChannel, _ := client.GetQRChannel(context.Background())
		err := client.Connect()

		if err != nil {
			log.Fatalln(err)
		}
		for event := range qrChannel {
			if event.Event == "code" {
				err = qrcode.WriteFile(event.Code, qrcode.Medium, 512, "qr.png")
				if err != nil {
					log.Println("Couldn't print qr because: ", err.Error())
				}
			} else {
				log.Println("Event: ", event.Event)
			}
		}
	} else {
		err := client.Connect()
		if err != nil {
			log.Fatalln(err)
		}
	}

	return client
}
