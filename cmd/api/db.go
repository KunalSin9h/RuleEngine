package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"time"
)

func setupPostgres(connString string) *sql.DB {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		slog.Error("Failed to connect to Postgres DB", "error message", err.Error())
		os.Exit(1)
	}

	slog.Info("Connected to Postgres DB...", "connection string", connString)

	backOff, retryCounter := 1, 1

	for {
		if err := db.Ping(); err != nil {
			slog.Info(fmt.Sprintf("Failed to connect, retrying... [%d/5]", retryCounter))

			if retryCounter >= 5 {
				break
			}

			retryCounter++
			backOff *= 2

			time.Sleep(time.Second * time.Duration(backOff))
		} else {
			slog.Info("Successfully connected to Postgres DB")
			break
		}
	}
	return db
}
