package main

import (
	"database/sql"
	"fmt"
	"gator/internal/database"
)

func mustOpenDB(dbURL string) (*database.Queries, *sql.DB) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.Ping()

	return database.New(db), db
}
