// Package ast provides an implementation of Abstract Syntax Tree for RuleEngine.
package ast

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// JSON is type helper that represent JSON like construct
type JSON map[string]any

type Node struct {
	Type  string `json:"type"`
	Value string `json:"value,omitempty"`
	Left  *Node  `json:"left,omitempty"`
	Right *Node  `json:"right,omitempty"`
	Field string `json:"field,omitempty"`
	Op    string `json:"op,omitempty"`
}

// EvaluateRule evaluates the rule against the provided data
func EvaluateRule(astJSON string, data JSON) (bool, error) {
	var root Node
	err := json.Unmarshal([]byte(astJSON), &root)
	if err != nil {
		return false, fmt.Errorf("failed to parse AST JSON: %v", err)
	}
	return evaluateNode(&root, data), nil
}

func evaluateNode(node *Node, data JSON) bool {
	switch node.Type {
	case "operator":
		return evaluateOperator(node, data)
	case "condition":
		return evaluateCondition(node, data)
	default:
		fmt.Printf("Unknown node type: %s\n", node.Type)
		return false
	}
}

func evaluateOperator(node *Node, data JSON) bool {
	leftResult := evaluateNode(node.Left, data)
	rightResult := evaluateNode(node.Right, data)

	switch node.Value {
	case "AND":
		return leftResult && rightResult
	case "OR":
		return leftResult || rightResult
	default:
		fmt.Printf("Unknown operator: %s\n", node.Value)
		return false
	}
}

func evaluateCondition(node *Node, data JSON) bool {
	fieldValue, ok := data[node.Field]
	if !ok {
		fmt.Printf("Field not found in data: %s\n", node.Field)
		return false
	}

	switch node.Op {
	case "=":
		return fmt.Sprintf("%v", fieldValue) == node.Value
	case "!=":
		return fmt.Sprintf("%v", fieldValue) != node.Value
	case ">", "<", ">=", "<=":
		return compareNumbers(fieldValue, node.Value, node.Op)
	default:
		fmt.Printf("Unknown comparison operator: %s\n", node.Op)
		return false
	}
}

func compareNumbers(fieldValue any, nodeValue, op string) bool {
	var fv, nv float64
	var err error

	// Convert fieldValue to float64
	switch v := fieldValue.(type) {
	case float64:
		fv = v
	case int:
		fv = float64(v)
	case string:
		fv, err = strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("Error converting field value to number: %v\n", err)
			return false
		}
	default:
		fmt.Printf("Unsupported field value type: %T\n", fieldValue)
		return false
	}

	// Convert nodeValue to float64
	nv, err = strconv.ParseFloat(nodeValue, 64)
	if err != nil {
		fmt.Printf("Error converting node value to number: %v\n", err)
		return false
	}

	switch op {
	case ">":
		return fv > nv
	case ">=":
		return fv >= nv
	case "<":
		return fv < nv
	case "<=":
		return fv <= nv
	default:
		fmt.Printf("Unexpected operator in compareNumbers: %s\n", op)
		return false
	}
}
