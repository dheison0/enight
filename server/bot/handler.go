package bot

import (
	"fmt"

	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(event any) {
	switch v := event.(type) {
	case *events.Message:
		if v.Info.IsFromMe || v.Info.IsGroup {
			return
		}
		fmt.Printf("New message from %s: %+v\n", v.Info.Sender.User, v.Message.GetConversation())
		fmt.Println(v.Info.PushName)
		markRead(v)
		err := sendText(v, "Oi, tudo bem? "+v.Info.PushName)
		if err != nil {
			fmt.Println("Deu pra enviar não oh!", err)
		} else {
			fmt.Println("Envou foi tudo!")
		}
	}
}
