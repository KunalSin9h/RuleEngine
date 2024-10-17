package tests

import (
	"encoding/json"
	"fmt"
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

func TestParser_CombineRules_AND(t *testing.T) {
	rules := []string{
		"(age >= 18 AND department = 'Marketing')",
		"(age >= 18 AND income >= 30000)",
	}

	node, err := parser.CombineRules(rules)
	if err != nil {
		t.Error(err.Error())
	}

	testCases := []TestCase{
		{
			expectedResult: true,
			data: ast.JSON{
				"age":        "20",
				"department": "Marketing",
				"income":     "30000",
			},
		},
		{
			expectedResult: true,
			data: ast.JSON{
				"age":        "30",
				"department": "Marketing",
				"income":     "90000",
			},
		},
		{
			expectedResult: false,
			data: ast.JSON{
				"age": "17",
			},
		},
		{
			expectedResult: false,
			data: ast.JSON{
				"age":        "27",
				"department": "Sales",
				"income":     "90000",
			},
		},
		{
			expectedResult: false,
			data: ast.JSON{
				"age":        "27",
				"department": "Marketing",
				"income":     "10000",
			},
		},
	}

	for _, testCase := range testCases {
		if node == nil {
			t.Error("Invalid AST Node")
			return
		}

		result := node.EvaluateNode(testCase.data)

		if result != testCase.expectedResult {
			t.Errorf("Expected result: %v, got %v for data: %v", testCase.expectedResult, result, testCase.data)
		}
	}
}

func TestParser_CombineRules_OR(t *testing.T) {
	rules := []string{
		"(age >= 18 OR department = 'Marketing')",
		"(age >= 18 OR income >= 30000)",
	}

	node, err := parser.CombineRules(rules)
	if err != nil {
		t.Error(err.Error())
	}

	d, _ := json.MarshalIndent(node, "", "  ")
	fmt.Println(string(d))
	//t.Fail()

	testCases := []TestCase{
		{
			expectedResult: true,
			data: ast.JSON{
				"age":        "40",
				"income":     "31000",
				"department": "R&D",
			},
		},
		{
			expectedResult: true,
			data: ast.JSON{
				"age":        "30",
				"income":     "10000",
				"department": "Marketing",
			},
		},
		{
			expectedResult: false,
			data: ast.JSON{
				"age":        "17",
				"income":     "10000",
				"department": "Marketing",
			},
		},
		{
			expectedResult: true,
			data: ast.JSON{
				"age":        "27",
				"department": "Sales",
				"income":     "10000",
			},
		},
	}

	for _, testCase := range testCases {
		if node == nil {
			t.Error("Invalid AST Node")
			return
		}

		result := node.EvaluateNode(testCase.data)

		if result != testCase.expectedResult {
			t.Errorf("Expected result: %v, got %v for data: %v", testCase.expectedResult, result, testCase.data)
		}
	}
}
