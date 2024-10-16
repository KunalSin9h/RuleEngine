package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

const (
	PORT = "5000"
	HOST = "0.0.0.0"
)

func main() {
	router := http.NewServeMux()

	// Show Frontend UI on `GET /`
	router.Handle("GET /", http.FileServer(http.Dir("./ui/dist/")))

	// Other Endpoints

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", HOST, PORT),
		Handler: router,
	}

	slog.Info("Starting server", "location", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Failed to start server", err)
	}
}
