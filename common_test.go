package common

import (
	"os"
	"testing"
	"time"
)

var now = time.Now().String()

var getEnvTestCases = []struct {
	Name          string
	VarName       string
	VarValue      string
	VarFallback   string
	ExpectedValue string
}{
	{
		Name:          "basic",
		VarName:       "BASIC_" + now,
		VarValue:      now,
		VarFallback:   "",
		ExpectedValue: now,
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
			os.Setenv(testCase.VarName, testCase.VarValue)

			// Test.
			actual := GetEnv(testCase.VarName, testCase.VarFallback)
			if actual != testCase.ExpectedValue {
				t.Errorf("got: %s, expected: %s", actual, testCase.ExpectedValue)
			}
		})
	}
}
