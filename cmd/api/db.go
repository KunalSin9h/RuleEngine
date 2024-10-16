package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

func setupPostgres(connString string) *sql.DB {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		slog.Error("Failed to connect to Postgres DB", "error message", err.Error())
		os.Exit(1)
	}

	return db
}
