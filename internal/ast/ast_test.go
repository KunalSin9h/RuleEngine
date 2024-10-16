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

func TestEvaluateRuleWithANDOperator(t *testing.T) {
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

func TestEvaluateRuleWithOROperator(t *testing.T) {
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
func TestEvaluateRuleGTEOperations(t *testing.T) {
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

func judge(t *testing.T, testCases []TestCase, astJSON string) {
	for _, testCase := range testCases {
		result, err := EvaluateRule(astJSON, testCase.data)

		if err != nil {
			t.Log(err.Error())
			t.Fail()
		}

		if result != testCase.expectedResult {
			t.Logf("Expected %v, received %v, for data %v", testCase.expectedResult, result, testCase.data)
			t.Fail()
		}
	}
}
