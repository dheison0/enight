package bot

import (
	"log"

	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(event any) {
	switch v := event.(type) {
	case *events.Message:
		if v.Info.IsFromMe || v.Info.IsGroup {
			return
		}
		log.Printf("New message from %s: %+v\n", v.Info.Sender.User, v.Message.GetConversation())
		log.Println(v.Info.PushName)
		markRead(v)
		err := sendText(v, "Oi, tudo bem? "+v.Info.PushName)
		if err != nil {
			log.Println("Deu pra enviar não oh!", err)
		} else {
			log.Println("Envou foi tudo!")
		}
	}
}
