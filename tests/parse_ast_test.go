package tests

import (
	"github.com/kunalsin9h/ruleengine/internal/ast"
	"github.com/kunalsin9h/ruleengine/internal/parser"
	"testing"
)

type TestCase struct {
	expectedResult bool
	data           ast.JSON
}

func TestParseAST(t *testing.T) {
	ruleString := `((age > 30 AND department = 'Sales') OR (age < 25 AND
		department = 'Marketing')) AND (salary > 50000 OR experience > 5)`

	astNode, err := parser.CreateRule(ruleString)

	// This should parse correctly
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	testCases := []TestCase{
		{
			expectedResult: true,
			data: ast.JSON{
				"age":        31,
				"department": "Sales",
				"salary":     80000,
			},
		},
		{
			expectedResult: true,
			data: ast.JSON{
				"age":        21,
				"department": "Marketing",
				"salary":     80000,
			},
		},
		{
			expectedResult: true,
			data: ast.JSON{
				"age":        21,
				"department": "Marketing",
				"experience": "6",
			},
		},
		{
			expectedResult: false,
			data: ast.JSON{
				"age":        24,
				"department": "Sales",
				"experience": "6",
			},
		},
		{
			expectedResult: false,
			data: ast.JSON{
				"age":        24,
				"department": "Marketing",
				"experience": "2",
			},
		},
		{
			expectedResult: false,
			data: ast.JSON{
				"age":        38,
				"department": "Sales",
				"experience": "4",
			},
		},
	}

	for _, testCase := range testCases {
		result := astNode.EvaluateNode(testCase.data)

		if result != testCase.expectedResult {
			t.Logf("Expected: %v, got %v for data %v", testCase.expectedResult, result, testCase.data)
			t.Error("Failed to evaluate AST Against data")
		}
	}
}
