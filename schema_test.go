package main

import (
	"database/sql"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestUserTableSchema(t *testing.T) {
	// 1. Connect to your local 'gator' database
	db, err := sql.Open("postgres", "postgres://postgres@localhost:5432/gator?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// 2. Query the information_schema to check column types
	rows, err := db.Query(`
		SELECT column_name, data_type, is_nullable
		FROM information_schema.columns
		WHERE table_name = 'users'
	`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	columns := make(map[string]string)
	for rows.Next() {
		var name, dtype, nullable string
		if err := rows.Scan(&name, &dtype, &nullable); err != nil {
			t.Fatal(err)
		}
		columns[name] = dtype
	}

	// 3. Assert requirements
	expected := map[string]string{
		"id":         "uuid",
		"created_at": "timestamp",
		"updated_at": "timestamp",
		"name":       "text",
	}

	for col, expectedType := range expected {
		actualType, ok := columns[col]
		if !ok {
			t.Errorf("expected collumn %s to exist, but it was not found", col)
			continue
		}

		if !strings.Contains(strings.ToLower(actualType), strings.ToLower(expectedType)) {
			t.Errorf("colum %s: expected type %s, got %s", col, expectedType, actualType)
		}
	}

}
