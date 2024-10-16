package main

import (
	"net/http"
)

func setupRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Show Frontend UI on `GET /`
	router.Handle("GET /", http.FileServer(http.Dir("./ui/")))

	// Other Endpoints

	return router
}
