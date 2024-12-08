package main

import (
	"os"
	"server/api"
	"server/bot"
	"server/database"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
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
