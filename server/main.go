package main

import (
	"os"
	"server/api"
	"server/bot"
	"server/database"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env", "../.env")

	debug := os.Getenv("DEBUG") == "true"
	addr := ":" + os.Getenv("PORT")
	webFiles := os.Getenv("WEB_FILES")
	dbPath := os.Getenv("DB_PATH")
	if webFiles == "" {
		panic("WEB_FILES environment variable not defined!\nIt's needed to show the webpage at /")
	} else if dbPath == "" {
		panic("DB_PATH not defined!")
	} else if err := database.Init(dbPath); err != nil { // initialize database and this if it's ok
		panic("failed to init database! " + err.Error())
	} else if addr == ":" {
		addr = ":8080"
	}

	bot.Start(debug)
	if err := api.Start(addr, webFiles, debug); err != nil {
		bot.Stop()
		panic(err)
	}
}
