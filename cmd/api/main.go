package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var (
	PORT     = "5000"
	HOST     = "0.0.0.0"
	POSTGRES = "postgresql://admin:admin@localhost:5432/rules?sslmode=disable"
)

func init() {
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	if os.Getenv("HOST") != "" {
		HOST = os.Getenv("HOST")
	}

	if os.Getenv("POSTGRES") != "" {
		POSTGRES = os.Getenv("POSTGRES")
	}
}

type Config struct {
	db     *sql.DB
	router *http.ServeMux
}

func main() {
	app := Config{}

	var err error
	app.db, err = setupPostgres(POSTGRES)

	if err != nil {
		os.Exit(1)
	}

	app.setupRouter()

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", HOST, PORT),
		Handler: app.router,
	}

	slog.Info("Starting server", "location", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Failed to start server", "error message", err.Error())
	}
}
