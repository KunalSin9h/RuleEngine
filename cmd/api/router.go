package main

import (
	"net/http"
)

func (c *Config) setupRouter() {
	router := http.NewServeMux()

	// Show Frontend UI on `GET /`
	router.Handle("GET /", http.FileServer(http.Dir("./ui/")))

	// Other Endpoints

	c.router = router
}
