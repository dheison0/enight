package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(path string) error {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	// db can be better configured!
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return errors.New("failed to enable foreign_keys! " + err.Error())
	}
	return err
}
