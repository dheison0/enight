package bot

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
  _ "modernc.org/sqlite"
)

var client *whatsmeow.Client

func Start(debug bool) {
  slog.Info("Starting bot...")
	// Initialize loggers
	logLevel := "WARN"
	if debug {
		logLevel = "DEBUG"
	}
	device, err := setupDatabase(logLevel).GetFirstDevice()
	if err != nil {
		panic("failed to get first device! " + err.Error())
	}
	client := setupWhatsappClient(logLevel, device)
	client.AddEventHandler(EventHandler)
	slog.Info("Bot started!")
}

func setupDatabase(logLevel string) *sqlstore.Container {
	dbLog := waLog.Stdout("Bot database", logLevel, false)
	botDBPath := os.Getenv("BOT_DB_PATH")
	if botDBPath == "" {
		botDBPath = "./bot.sqlite3"
		slog.Warn("BOT_DB_PATH not provided, using default", slog.String("botDBPath", botDBPath))
	}
	container, err := sqlstore.New("sqlite", fmt.Sprintf("file:%s?_pragma=foreign_keys(1)", botDBPath), dbLog)
	if err != nil {
		panic("failed to create bot database container! " + err.Error())
	}
	return container
}

func setupWhatsappClient(logLevel string, device *store.Device) *whatsmeow.Client {
	botLog := waLog.Stdout("Bot", logLevel, false)
	client = whatsmeow.NewClient(device, botLog)
	if client.Store.ID == nil {
		botLog.Warnf("No login detected! Starting QR flow...")
		qrChan, _ := client.GetQRChannel(context.Background())
		err := client.Connect()
		if err != nil {
			panic("failed to connect bot! " + err.Error())
		}
		for event := range qrChan {
			if event.Event == "code" {
				botLog.Warnf("New QR code received")
				qrterminal.GenerateHalfBlock(event.Code, qrterminal.L, os.Stdout)
			} else {
				botLog.Warnf("Event received: %s", event.Event)
			}
		}
	} else {
		err := client.Connect()
		if err != nil {
			panic("failed to connect bot! " + err.Error())
		}
	}
	return client
}

func Stop() {
	slog.Info("Stopping bot...")
	client.Disconnect()
}
