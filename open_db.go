package main

import (
	"database/sql"
	"errors"
	"gator/internal/database"
)

func openDB(dbURL string) (*database.Queries, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return database.New(db), errors.New("not enough arguments")
	}
	//defer db.Close()
	//db.Ping()

	return database.New(db), nil
}
