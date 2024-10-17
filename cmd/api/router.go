package main

import (
	"net/http"
)

func (c *Config) setupRouter() {
	router := http.NewServeMux()

	// Show Frontend UI on `GET /`
	router.Handle("GET /", http.FileServer(http.Dir("./ui/")))

	// Get all rules in the engine with GET /rules
	router.HandleFunc("GET /rules", c.getAllRules)

	// Get single rule with id
	router.HandleFunc("GET /rule/{id}", func(writer http.ResponseWriter, request *http.Request) {})
	// Create a new rule
	router.HandleFunc("POST /rule", func(w http.ResponseWriter, r *http.Request) {})
	// Update existing rule
	router.HandleFunc("PATCH /rule/{id}", func(writer http.ResponseWriter, request *http.Request) {})
	// Delete a rule
	router.HandleFunc("DELETE /rule/{id}", func(writer http.ResponseWriter, request *http.Request) {})

	// Merge multiple rules
	router.HandleFunc("POST /rules/merge", func(w http.ResponseWriter, r *http.Request) {})

	c.router = router
}

func (c *Config) getAllRules(w http.ResponseWriter, r *http.Request) {
}
