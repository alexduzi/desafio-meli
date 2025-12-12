package database

import (
	"fmt"
	"project/internal/infra/database/migrations"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	schema, err := migrations.FS.ReadFile("001_schema.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read schema: %w", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return nil, fmt.Errorf("failed to execute schema: %w", err)
	}

	seed, err := migrations.FS.ReadFile("002_seed.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read seed data: %w", err)
	}

	if _, err := db.Exec(string(seed)); err != nil {
		return nil, fmt.Errorf("failed to execute seed data: %w", err)
	}

	return db, nil
}
