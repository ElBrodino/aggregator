package main

import (
	"database/sql"
	"errors"
	"gator/internal/database"
)

func openDB(dbURL string) (*database.Queries, *sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return database.New(db), db, errors.New("not enough arguments")
	}
	defer db.Close()
	db.Ping()

	return database.New(db), db, nil
}
