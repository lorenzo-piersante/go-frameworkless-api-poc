package storage

import (
	"database/sql"
	"log"
	"os"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	err := executeMigrations(db)
	if err != nil {
		log.Printf("Error occurred during migrations execution: %v", err)
		os.Exit(1)
	}

	return &Storage{db: db}
}

func executeMigrations(db *sql.DB) error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id STRING PRIMARY KEY, username STRING, password STRING)")
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}
