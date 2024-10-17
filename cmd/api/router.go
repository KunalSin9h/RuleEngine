package main

import (
	"encoding/json"
	"fmt"
	"github.com/kunalsin9h/ruleengine/internal/parser"
	"io"
	"net/http"
)

func (c *Config) setupRouter() {
	router := http.NewServeMux()

	// Show Frontend UI on `GET /`
	router.Handle("GET /", http.FileServer(http.Dir("./ui/dist/")))

	// CREATE RULE
	// Create a new rule
	router.HandleFunc("POST /rule", enableCORS(c.createRule))

	// COMBINE RULES
	router.HandleFunc("POST /rules", func(w http.ResponseWriter, r *http.Request) {})

	// EVALUATE RULE
	router.HandleFunc("GET /rule/eval", func(w http.ResponseWriter, r *http.Request) {})

	// OTHER HELPER ENDPOINTS

	// GET ALL RULES
	router.HandleFunc("GET /rules", func(w http.ResponseWriter, r *http.Request) {})

	c.router = router
}

type CreateRulePayload struct {
	Rule        string `json:"rule"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Config) createRule(w http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer r.Body.Close()

	var payload CreateRulePayload
	err = json.Unmarshal(data, &payload)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	rule := payload.Rule
	/*	name := payload.Name
		description := payload.Description*/

	ast, err := parser.CreateRule(rule)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	responseData, err := json.MarshalIndent(ast, "", "\t")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(responseData))
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}
