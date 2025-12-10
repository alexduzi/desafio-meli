package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	schema, err := os.ReadFile("migrations/001_schema.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read schema: %w", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return nil, fmt.Errorf("failed to execute schema: %w", err)
	}

	seed, err := os.ReadFile("migrations/002_seed.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read seed data: %w", err)
	}

	if _, err := db.Exec(string(seed)); err != nil {
		return nil, fmt.Errorf("failed to execute seed data: %w", err)
	}

	return db, nil
}
