package bot

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var client *whatsmeow.Client

func Start(debug bool) {
	log.Println("Starting bot...")
	// Initialize loggers
	logLevel := "WARN"
	if debug {
		logLevel = "DEBUG"
	}
	botLog := waLog.Stdout("Bot", logLevel, false)
	dbLog := waLog.Stdout("Bot database", logLevel, false)

	// Connect to database that will store the session
	botDBPath := os.Getenv("BOT_DB_PATH")
	container, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", botDBPath), dbLog)
	if err != nil {
		panic("failed to create bot database container! " + err.Error())
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic("failed to get first device! " + err.Error())
	}

	// Connect to WhatsApp
	client = whatsmeow.NewClient(deviceStore, botLog)
	if client.Store.ID == nil {
		// No login detected
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic("failed to connect bot! " + err.Error())
		}
		for event := range qrChan {
			if event.Event == "code" {
				botLog.Infof("New QR code received")
				qrterminal.GenerateHalfBlock(event.Code, qrterminal.L, os.Stdout)
			} else {
				botLog.Infof("Event received: %s", event.Event)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			panic("failed to connect bot! " + err.Error())
		}
	}
	client.AddEventHandler(EventHandler)
	log.Println("Bot started!")
}

func Stop() {
	log.Println("Stopping bot...")
	client.Disconnect()
}
