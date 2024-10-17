package main

import (
	"net/http"
)

func (c *Config) setupRouter() {
	router := http.NewServeMux()

	// Show Frontend UI on `GET /`
	router.Handle("GET /", http.FileServer(http.Dir("./ui/")))

	// CREATE RULE
	// Create a new rule
	router.HandleFunc("POST /rule", func(w http.ResponseWriter, r *http.Request) {})

	// COMBINE RULES
	router.HandleFunc("POST /rules", func(w http.ResponseWriter, r *http.Request) {})

	// EVALUATE RULE
	router.HandleFunc("GET /rule/eval", func(w http.ResponseWriter, r *http.Request) {})

	// OTHER HELPER ENDPOINTS

	// GET ALL RULES
	router.HandleFunc("GET /rules", func(w http.ResponseWriter, r *http.Request) {})

	c.router = router
}
