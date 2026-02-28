package main

import (
	"gator/internal/database"
	"os"
)

func mustOpenDB(dbURL string) *database.Queries {
	dbQueries, err := openDB(dbURL)
	if err != nil {
		os.Exit(1)
	}

	return dbQueries
}
