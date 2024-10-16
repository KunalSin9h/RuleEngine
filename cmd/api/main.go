package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
)

const (
	PORT     = "5000"
	HOST     = "0.0.0.0"
	POSTGRES = "postgresql://admin:admin@localhost:5432/rules"
)

type Config struct {
	db     *sql.DB
	router *http.ServeMux
}

func main() {
	app := Config{
		db: setupPostgres(POSTGRES),
	}

	app.setupRouter()

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", HOST, PORT),
		Handler: app.router,
	}

	slog.Info("Starting server", "location", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Failed to start server", err)
	}
}
