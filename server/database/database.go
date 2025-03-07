package database

import (
	"crypto/sha256"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"os"

	_ "modernc.org/sqlite"
)

const (
	DEFAULT_PASSWORD       = "admin"
	DEFAULT_SHIPPING_PRICE = 1
)

var db *sql.DB

//go:embed migrations/*.sql
var migrations embed.FS

func Init() {
	var err error
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "./system.sqlite3"
		slog.Warn("DB_PATH not provided, using default", slog.String("systemDBPath", path))
	}
	db, err = sql.Open("sqlite", path)
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
		slog.Info("Running migration", slog.String("file", file.Name()))
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
	settings, err := GetSettings()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if settings.PasswordHash == "" {
		slog.Info("Inserting settings into database...")
		hash := sha256.Sum256([]byte(DEFAULT_PASSWORD))
		_, err = db.Exec(
			"INSERT INTO settings(shipping_price, password_hash) VALUES(?, ?);",
			DEFAULT_SHIPPING_PRICE, fmt.Sprintf("%x", hash),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close closes the database connection. It must be called before the program
// exits.
func Close() {
	db.Close()
}
