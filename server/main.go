package main

import (
	"log"
	"os"
	"server/api"
	"server/bot"
	"server/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not loaded")
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
