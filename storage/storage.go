package storage

import (
	"context"
	"database/sql"
	"log"
	"os"
)

var ctx = context.Background()

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	err := executeMigrations(db)
	if err != nil {
		log.Printf("Error occurred during migrations execution: %v", err)
		os.Exit(1)
	}

	err = executeMigrations(db)

	return &Storage{db: db}
}

func executeMigrations(db *sql.DB) error {
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id STRING PRIMARY KEY, username STRING, password STRING)")
	_, _ = statement.Exec()
	return
}
