package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kunalsin9h/ruleengine/internal/ast"
	"github.com/kunalsin9h/ruleengine/internal/db"
	"github.com/kunalsin9h/ruleengine/internal/parser"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func (c *Config) setupRouter() {
	router := http.NewServeMux()

	// Show Frontend UI on `GET /`
	router.Handle("GET /", http.FileServer(http.Dir("./ui/dist/")))

	// CREATE RULE
	// Create a new rule
	router.HandleFunc("POST /rule", enableCORS(c.createRule))

	// COMBINE RULES
	router.HandleFunc("POST /rules", enableCORS(c.combineRules))

	// EVALUATE RULE
	router.HandleFunc("POST /rule/eval", enableCORS(c.evalRule))

	// OTHER HELPER ENDPOINTS

	// GET ALL RULES
	router.HandleFunc("GET /rules", enableCORS(c.getAllRules))

	c.router = router
}

type CreateRulePayload struct {
	Rule        string `json:"rule"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// createRule is an endpoint handler for creating new rule with rule string
func (c *Config) createRule(w http.ResponseWriter, r *http.Request) {
	// ready rule string in the request body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(&w, err)
		return
	}
	defer r.Body.Close()

	var payload CreateRulePayload
	err = json.Unmarshal(data, &payload)

	if err != nil {
		sendError(&w, err)
		return
	}

	rule := payload.Rule
	name := payload.Name
	description := payload.Description

	// parse the AST with the given rule string
	ast, err := parser.CreateRule(rule)

	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	// convert the AST into JSON format
	astJson, err := json.MarshalIndent(ast, "", "\t")

	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	// Add the content in the database
	query := db.New(c.db)
	err = query.CreateRule(context.Background(), db.CreateRuleParams{
		Name:        name,
		Description: pgtype.Text{String: description, Valid: true},
		Rule:        rule,
		Ast:         astJson,
	})

	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	// Send the AST JSON Representation
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(astJson))
}

func (c *Config) getAllRules(w http.ResponseWriter, r *http.Request) {
	query := db.New(c.db)

	rules, err := query.GetRules(context.Background())
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	data, err := json.MarshalIndent(rules, "", "\t")

	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(data))
}

type EvalRulePayload struct {
	RuleID int32  `json:"rule_id"`
	User   string `json:"user"`
}

func (c *Config) evalRule(w http.ResponseWriter, r *http.Request) {
	// ready rule id and user json in the request body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(&w, err)
		return
	}
	defer r.Body.Close()

	var payload EvalRulePayload
	err = json.Unmarshal(data, &payload)

	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	// fetch AST Node with ID = payload.RuleID
	query := db.New(c.db)
	rule, err := query.GetRule(context.Background(), payload.RuleID)

	var astNode ast.Node
	err = json.Unmarshal(rule.Ast, &astNode)

	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	var userDataJson map[string]any
	err = json.Unmarshal([]byte(payload.User), &userDataJson)
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	result := astNode.EvaluateNode(userDataJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responseData, err := json.Marshal(map[string]any{
		"result": result,
	})
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(responseData))
}

type CombineRulesPayload struct {
	Rules       []string `json:"rules"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
}

func (c *Config) combineRules(w http.ResponseWriter, r *http.Request) {
	// Ready request body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(&w, err)
		return
	}

	var payload CombineRulesPayload
	err = json.Unmarshal(data, &payload)
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	astNode, err := parser.CombineRules(payload.Rules)
	if err != nil {
		slog.Error("Failed to parse")
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	// convert astNode into JSON
	astJson, err := json.MarshalIndent(astNode, "", "\t")
	if err != nil {
		slog.Error("Failed to marshal")
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	// we have successfully parsed, not add them to database
	query := db.New(c.db)
	err = query.CreateRule(context.Background(), db.CreateRuleParams{
		Name:        payload.Name,
		Description: pgtype.Text{String: payload.Description, Valid: true},
		Ast:         astJson,
		Rule:        strings.Join(payload.Rules, "; "),
	})

	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(astJson))
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

func sendError(w *http.ResponseWriter, err error, code ...int) {
	slog.Error(err.Error())
	(*w).WriteHeader(http.StatusBadRequest)

	if len(code) > 0 {
		(*w).WriteHeader(code[0])
	}

	(*w).Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(*w, err.Error())
}
