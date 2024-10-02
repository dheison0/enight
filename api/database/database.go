package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(path string) error {
	var err error
	db, err = sql.Open("sqlite3", path)
	// db can be better configured!
	return err
}
