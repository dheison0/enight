package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	var err error
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "./system.sqlite3"
		log.Println("DB_PATH not provided, using ./system.sqlite3")
	}
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		panic("Failed to open database! " + err.Error())
	}
	// TODO: db can be better configured!
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic("failed to enable foreign_keys! " + err.Error())
	}
}

// Close closes the database connection. It must be called before the program
// exits.
func Close() {
	db.Close()
}
