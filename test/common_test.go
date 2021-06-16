package test

import (
	"github.com/foresthoffman/common"
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
			actual := common.GetEnv(testCase.VarName, testCase.VarFallback)
			if actual != testCase.ExpectedValue {
				t.Errorf("got: %s, expected: %s", actual, testCase.ExpectedValue)
			}
		})
	}
}

var urlTestCases = []struct {
	Name     string
	Paths    []string
	Expected string
}{
	{
		Name:     "single-empty",
		Paths:    []string{""},
		Expected: "/",
	},
	{
		Name:     "multi-empty",
		Paths:    []string{" ", ""},
		Expected: "/",
	},
	{
		Name:     "basic-relative",
		Paths:    []string{"tmp"},
		Expected: "tmp",
	},
	{
		Name:     "basic-absolute",
		Paths:    []string{"/tmp"},
		Expected: "/tmp",
	},
	{
		Name:     "mutli-relative",
		Paths:    []string{"tmp", "test"},
		Expected: "tmp/test",
	},
	{
		Name:     "multi-absolute",
		Paths:    []string{"/tmp", "test"},
		Expected: "/tmp/test",
	},
	{
		Name:     "slash",
		Paths:    []string{"/"},
		Expected: "/",
	},
	{
		Name:     "slash-2",
		Paths:    []string{"a", "/"},
		Expected: "a/",
	},
	{
		Name:     "slashed-relative",
		Paths:    []string{"tmp", "/test"},
		Expected: "tmp/test",
	},
	{
		Name:     "slashed-relative-2",
		Paths:    []string{"tmp/", "/test/"},
		Expected: "tmp/test",
	},
	{
		Name:     "slashed-relative-3",
		Paths:    []string{"tmp//", "/ /test/", "/"},
		Expected: "tmp/test/",
	},
	{
		Name:     "slashed-absolute",
		Paths:    []string{"/tmp", "/test"},
		Expected: "/tmp/test",
	},
	{
		Name:     "slashed-absolute-2",
		Paths:    []string{"/tmp/", "/test/"},
		Expected: "/tmp/test",
	},
	{
		Name:     "slashed-absolute-3",
		Paths:    []string{"///tmp/", "/test/"},
		Expected: "/tmp/test",
	},
	{
		Name:     "http",
		Paths:    []string{"http://localhost:9001", "api/v1", "agents", "1234abcdefg"},
		Expected: "http://localhost:9001/api/v1/agents/1234abcdefg",
	},
	{
		Name:     "single-router",
		Paths:    []string{"//agents/"},
		Expected: "/agents",
	},
	{
		Name:     "multi-router",
		Paths:    []string{"/", "/agents/"},
		Expected: "/agents",
	},
	{
		Name:     "multi-router-2",
		Paths:    []string{"/", "/agents/", "/"},
		Expected: "/agents/",
	},
}

func TestUrl(t *testing.T) {
	for _, testCase := range urlTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := common.Url(testCase.Paths...)
			if actual != testCase.Expected {
				t.Errorf("got: %q, expected: %q", actual, testCase.Expected)
				return
			}
		})
	}
}

var pathTestCases = []struct {
	Name     string
	Paths    []string
	Expected string
}{
	{
		Name:     "single-empty",
		Paths:    []string{""},
		Expected: "/",
	},
	{
		Name:     "multi-empty",
		Paths:    []string{" ", ""},
		Expected: "/",
	},
	{
		Name:     "basic-relative",
		Paths:    []string{"tmp"},
		Expected: "tmp",
	},
	{
		Name:     "basic-absolute",
		Paths:    []string{"/tmp"},
		Expected: "/tmp",
	},
	{
		Name:     "mutli-relative",
		Paths:    []string{"tmp", "test"},
		Expected: "tmp/test",
	},
	{
		Name:     "multi-absolute",
		Paths:    []string{"/tmp", "test"},
		Expected: "/tmp/test",
	},
	{
		Name:     "slash",
		Paths:    []string{"/"},
		Expected: "/",
	},
	{
		Name:     "slash-2",
		Paths:    []string{"a", "/"},
		Expected: "a/",
	},
	{
		Name:     "slashed-relative",
		Paths:    []string{"tmp", "/test"},
		Expected: "tmp/test",
	},
	{
		Name:     "slashed-relative-2",
		Paths:    []string{"tmp/", "/test/"},
		Expected: "tmp/test",
	},
	{
		Name:     "slashed-relative-3",
		Paths:    []string{"tmp//", "/ /test/", "/"},
		Expected: "tmp/test/",
	},
	{
		Name:     "slashed-absolute",
		Paths:    []string{"/tmp", "/test"},
		Expected: "/tmp/test",
	},
	{
		Name:     "slashed-absolute-2",
		Paths:    []string{"/tmp/", "/test/"},
		Expected: "/tmp/test",
	},
	{
		Name:     "slashed-absolute-3",
		Paths:    []string{"///tmp/", "/test/"},
		Expected: "/tmp/test",
	},
	{
		Name:     "http",
		Paths:    []string{"http://localhost:9001", "api/v1", "agents", "1234abcdefg"},
		Expected: "http://localhost:9001/api/v1/agents/1234abcdefg",
	},
	{
		Name:     "single-router",
		Paths:    []string{"//agents/"},
		Expected: "/agents",
	},
	{
		Name:     "multi-router",
		Paths:    []string{"/", "/agents/"},
		Expected: "/agents",
	},
	{
		Name:     "multi-router-2",
		Paths:    []string{"/", "/agents/", "/"},
		Expected: "/agents/",
	},
}

func TestPath(t *testing.T) {
	for _, testCase := range pathTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := common.Path("/", testCase.Paths...)
			if actual != testCase.Expected {
				t.Errorf("got: %q, expected: %q", actual, testCase.Expected)
				return
			}
		})
	}
}

var dashTestCases = []struct {
	Name     string
	Paths    []string
	Expected string
}{
	{
		Name:     "single-empty",
		Paths:    []string{""},
		Expected: "",
	},
	{
		Name:     "multi-empty",
		Paths:    []string{"", "   "},
		Expected: "",
	},
	{
		Name:     "multi-empty-dashed",
		Paths:    []string{"--", "  - "},
		Expected: "",
	},
	{
		Name:     "single-suffix",
		Paths:    []string{"foo", "bar"},
		Expected: "foo-bar",
	},
	{
		Name:     "single-suffix-empty",
		Paths:    []string{"foo", ""},
		Expected: "foo",
	},
	{
		Name:     "single-suffix-dashed",
		Paths:    []string{"- -foo-", "- - bar-"},
		Expected: "foo-bar",
	},
	{
		Name:     "multi-suffix",
		Paths:    []string{"foo", "bar", "baz"},
		Expected: "foo-bar-baz",
	},
	{
		Name:     "multi-suffix-dashed",
		Paths:    []string{"-foo--", "--bar - -", " - baz -- "},
		Expected: "foo-bar-baz",
	},
	{
		Name:     "multi-suffix-middle-empty",
		Paths:    []string{"foo", "", "baz"},
		Expected: "foo-baz",
	},
	{
		Name:     "multi-suffix-middle-empty-dashed",
		Paths:    []string{"foo", "- - ---- ", "baz"},
		Expected: "foo-baz",
	},
}

func TestDash(t *testing.T) {
	for _, testCase := range dashTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := common.Dash(testCase.Paths...)
			if actual != testCase.Expected {
				t.Errorf("got: %q, expected: %q", actual, testCase.Expected)
				return
			}
		})
	}
}
