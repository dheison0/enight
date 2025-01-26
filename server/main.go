package main

import (
	"log/slog"
	"os"
	"server/api"
	"server/bot"
	"server/database"

	"github.com/joho/godotenv"
)

func init() {
  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
  slog.SetDefault(logger)
}

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Info(".env file not loaded")
	}
	debug := os.Getenv("DEBUG") == "true"

	database.Init()
	defer database.Close()
	if os.Getenv("DISABLE_BOT") != "true" {
		bot.Start(debug)
		defer bot.Stop()
	}
	if err := api.Start(debug); err != nil {
		panic(err)
	}
}
