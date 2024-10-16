package parser

import (
	"testing"
)

var parser Parser

func TestParser_Parse_Example1(t *testing.T) {
	ruleString := `((age > 30 AND department = 'Sales') OR (age < 25 AND
		department = 'Marketing')) AND (salary > 50000 OR experience > 5)`

	_, err := parser.Parse(ruleString)

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestParser_Parse_Example2(t *testing.T) {
	ruleString := `((age > 30 AND department = 'Marketing')) AND (salary >
		20000 OR experience > 5)`

	_, err := parser.Parse(ruleString)

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestParser_Parse_InvalidRule_MissingClosingParenthesis(t *testing.T) {
	ruleString := `((age > 30 AND department = 'Marketing')) AND (salary >d
		20000 OR experience > 5` // missing closing `)`

	_, err := parser.Parse(ruleString)

	if err == nil {
		t.Error("Parser works on invalid rule string")
		t.Fail()
	}
}

func TestParser_Parse_InvalidRule_MissingOperator(t *testing.T) {
	// missing AND operator between two expressions
	//                       * here
	ruleString := `((age > 30 department = 'Marketing')) AND (salary >d
		20000 OR experience > 5)`

	_, err := parser.Parse(ruleString)

	if err == nil {
		t.Error("Parser works on invalid rule string")
	}
}

func TestParser_Parse_InvalidRule_BadComparison(t *testing.T) {
	ruleString := `((age > 30 department = 'Marketing')) AND (salary >
		20000 OR experience > abcd)` // * here

	_, err := parser.Parse(ruleString)

	// Should fail
	if err == nil {
		t.Error("Parser works on invalid rule string")
	}
}
