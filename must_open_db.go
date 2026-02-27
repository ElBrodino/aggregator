package main

import (
	"database/sql"
	"gator/internal/database"
	"os"
)

func mustOpenDB(dbURL string) (*database.Queries, *sql.DB) {
	dbQueries, db, err := openDB(dbURL)
	if err != nil {
		os.Exit(1)
	}

	return dbQueries, db
}
