package ast

import (
	"strings"
	"testing"
)

type TestCase struct {
	expectedResult bool
	data           JSON
}

var astJSON = `{
		"type": "operator",
		"value": "AND",
		"left": {
			"type": "condition",
			"field": "age",
			"op": ">",
			"value": "30"
		},
		"right": {
			"type": "condition",
			"field": "department",
			"op": "=",
			"value": "Sales"
		}
	}`

func TestEvaluateRule_WithANDOperator(t *testing.T) {
	testCases := []TestCase{
		{
			expectedResult: true,
			data: JSON{
				"age":        35,
				"department": "Sales",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"age":        15,
				"department": "Sales",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"age":        45,
				"department": "R&D",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"age":        15,
				"department": "R&D",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"income": "65000",
			},
		},
	}

	judge(t, testCases, astJSON)
}

func TestEvaluateRule_WithOROperator(t *testing.T) {
	ast := strings.ReplaceAll(astJSON, "AND", "OR")

	testCases := []TestCase{
		{
			expectedResult: true,
			data: JSON{
				"age":        35,
				"department": "Sales",
			},
		},
		{
			expectedResult: true,
			data: JSON{
				"age":        15,
				"department": "Sales",
			},
		},
		{
			expectedResult: true,
			data: JSON{
				"age":        45,
				"department": "R&D",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"age":        15,
				"department": "R&D",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"income": "65000",
			},
		},
	}

	judge(t, testCases, ast)
}

// Greater then Equal
func TestEvaluateRule_GTEOperations(t *testing.T) {
	ast := `{
		"type": "condition",
		"field": "income",
		"op": ">=",
		"value": "30000"
	}`

	testCases := []TestCase{
		{
			expectedResult: true,
			data: JSON{
				"income": "30000",
			},
		},
		{
			expectedResult: true,
			data: JSON{
				"income": "60000",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"income": "28000",
			},
		},
	}

	judge(t, testCases, ast)
}

// Less than Equal
func TestEvaluateRule_LTEOperations(t *testing.T) {
	ast := `{
		"type": "condition",
		"field": "income",
		"op": "<=",
		"value": "100000"
	}`

	testCases := []TestCase{
		{
			expectedResult: true,
			data: JSON{
				"income": "30000",
			},
		},
		{
			expectedResult: true,
			data: JSON{
				"income": "100000",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"income": "128000",
			},
		},
	}

	judge(t, testCases, ast)
}

// Not equal
func TestEvaluateRule_NotEqOperations(t *testing.T) {
	ast := `{
		"type": "condition",
		"field": "age",
		"op": "!=",
		"value": "10"
	}`

	testCases := []TestCase{
		{
			expectedResult: true,
			data: JSON{
				"age": "11",
			},
		},
		{
			expectedResult: false,
			data: JSON{
				"age": "10",
			},
		},
	}

	judge(t, testCases, ast)
}

func judge(t *testing.T, testCases []TestCase, astJSON string) {
	for _, testCase := range testCases {
		result, err := EvaluateRule(astJSON, testCase.data)

		if err != nil {
			t.Error(err)
		}

		if result != testCase.expectedResult {
			t.Errorf("Expected %v, received %v, for data %v", testCase.expectedResult, result, testCase.data)
		}
	}
}
