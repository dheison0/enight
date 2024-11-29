package database

import (
	"database/sql"
	"embed"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//go:embed migrations/*.sql
var migrations embed.FS

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
	if err = runMigrations(); err != nil {
		panic("failed to run migrations! " + err.Error())
	}
	if err = insertInitialValues(); err != nil {
		panic("failed to insert initial values! " + err.Error())
	}
}

func runMigrations() error {
	files, err := migrations.ReadDir("migrations")
	if err != nil {
		return errors.New("failed to read migrations directory! " + err.Error())
	}
	for _, file := range files {
		log.Println("Running migration:", file.Name())
		data, err := migrations.ReadFile("migrations/" + file.Name())
		if err != nil {
			return errors.New("can't read migration file! " + err.Error())
		}
		_, err = db.Exec(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func insertInitialValues() error {
	// TODO: insert initial values
	return nil
}

// Close closes the database connection. It must be called before the program
// exits.
func Close() {
	db.Close()
}
