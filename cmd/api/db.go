package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

func setupPostgres(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		slog.Error("Failed to connect to Postgres DB", "error message", err.Error())
		return nil, err
	}
	defer db.Close()

	slog.Info("Connecting to Postgres DB...", "connection string", connString)

	backOff, retryCounter := 1, 1

	for {
		if err := db.Ping(); err != nil {
			slog.Info(fmt.Sprintf("Failed to connect, retrying... [%d/5]", retryCounter))

			if retryCounter >= 5 {
				slog.Error("Exiting...")
				return nil, fmt.Errorf("error failed connection")
			}

			retryCounter++
			backOff *= 2

			time.Sleep(time.Second * time.Duration(backOff))
		} else {
			slog.Info("Successfully connected to Postgres DB")
			break
		}
	}

	return db, nil
}
