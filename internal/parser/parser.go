// Package parser provides way to parse "Rule String" into "AST"
package parser

import (
	"fmt"
	"github.com/kunalsin9h/ruleengine/internal/ast"
	"strings"
)

// Parser is a type for maintaining Parsing State
type Parser struct {
	tokens []string
	pos    int
}

// CreateRule  creates AST Node with ruleString
func CreateRule(ruleString string) (*ast.Node, error) {
	p := Parser{}
	return p.Parse(ruleString)
}

func CombineRules(ruleStrings []string) (*ast.Node, error) {
	trees := make([]*ast.Node, len(ruleStrings))

	for i, rs := range ruleStrings {
		tree, err := CreateRule(rs)
		if err != nil {
			return nil, err
		}

		trees[i] = tree
	}

	combineRoot := &ast.Node{
		Type:  "operator",
		Value: "AND",
	}

	current := combineRoot

	for _, tree := range trees {
		if current.Left == nil {
			current.Left = tree
		} else if current.Right == nil {
			current.Right = tree
		} else {
			newNode := &ast.Node{
				Left:  current.Right,
				Type:  "operator",
				Value: "AND",
			}

			current.Right = newNode
			current = newNode
		}
	}

	// Optimize AST to remove redundant checks.
	return ast.Optimize(combineRoot), nil
}

// Parse creates AST Node with rule string
func (p *Parser) Parse(ruleString string) (*ast.Node, error) {
	ruleString = strings.ReplaceAll(ruleString, "(", " ( ")
	ruleString = strings.ReplaceAll(ruleString, ")", " ) ")

	p.tokens = strings.Fields(ruleString)
	p.pos = 0

	return p.parseExpression()
}

func (p *Parser) parseExpression() (*ast.Node, error) {
	node, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for p.pos < len(p.tokens) && (p.tokens[p.pos] == "AND" || p.tokens[p.pos] == "OR") {
		op := p.tokens[p.pos]
		p.pos++

		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}

		node = &ast.Node{
			Type:  "operator",
			Value: op, // AND/OR
			Left:  node,
			Right: right,
		}
	}

	return node, nil
}

func (p *Parser) parseTerm() (*ast.Node, error) {
	if p.tokens[p.pos] == "(" {
		p.pos++
		node, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			return nil, fmt.Errorf("expected closing parenthesis")
		}
		p.pos++
		return node, nil
	}

	return p.parseCondition()
}

func (p *Parser) parseCondition() (*ast.Node, error) {
	if p.pos+3 >= len(p.tokens) {
		return nil, fmt.Errorf("unexpected end of input")
	}

	field := p.tokens[p.pos]
	p.pos++

	op := p.tokens[p.pos]
	p.pos++

	value := p.tokens[p.pos]
	p.pos++

	// Remove quotes from string values
	value = strings.Trim(value, "'")

	return &ast.Node{
		Type:  "condition",
		Field: field,
		Op:    op,
		Value: value,
	}, nil
}
