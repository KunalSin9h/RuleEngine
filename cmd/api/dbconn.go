package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"os"
	"time"
)

func setupPostgres(ctx context.Context, connString string) (*pgx.Conn, error) {
	slog.Info("Connecting to Postgres DB...", "connection string", connString)

	for i := 1; i <= 5; i++ {
		conn, err := pgx.Connect(ctx, connString)

		if err != nil {
			slog.Warn(fmt.Sprintf("Failed to connect to Postgres DB, retrying...[%d/5]", i))

			backOff := i * 2
			time.Sleep(time.Duration(backOff) * time.Second)
			continue
		}

		slog.Info("Successfully connected to Postgres DB")

		schema, err := os.ReadFile("./schema/schema.sql")
		if err != nil {
			slog.Error("Failed to read schema file")
			return nil, err
		}

		sqlSchema := string(schema)
		_, err = conn.Exec(ctx, sqlSchema)

		if err != nil {
			slog.Error("Failed to execute schema statement")
			return nil, err
		}

		return conn, err
	}

	slog.Error("Failed to connect to Postgres DB")
	return nil, fmt.Errorf("failed to connect to postgres db, exiting")
}
