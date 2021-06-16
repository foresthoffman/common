package common

import (
	"os"
	"testing"
)

var getEnvTestCases = []struct {
	Name          string
	VarName       string
	VarValue      string
	VarFallback   string
	ExpectedValue string
}{
	{
		Name:          "basic",
		VarName:       "BASIC_123",
		VarValue:      "123",
		VarFallback:   "",
		ExpectedValue: "123",
	},
	{
		Name:          "fallback",
		VarName:       "SUPERCALIFRAGILISTICEXPIALIDOCIOUS",
		VarValue:      "",
		VarFallback:   "Mary Poppins",
		ExpectedValue: "Mary Poppins",
	},
}

func TestGetEnv(t *testing.T) {
	for _, testCase := range getEnvTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			// Setup.
			err := os.Setenv(testCase.VarName, testCase.VarValue)
			if err != nil {
				t.Error(err)
				return
			}

			// Test.
			actual := GetEnv(testCase.VarName, testCase.VarFallback)
			if actual != testCase.ExpectedValue {
				t.Errorf("got: %s, expected: %s", actual, testCase.ExpectedValue)
			}
		})
	}
}
