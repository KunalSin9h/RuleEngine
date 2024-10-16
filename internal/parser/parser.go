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
