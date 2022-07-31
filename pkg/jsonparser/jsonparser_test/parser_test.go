package jsonparser_test

import (
	"reflect"
	"strings"
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func TestParser(tester *testing.T) {
	js := []string{
		`{`,
		`    "jsonparser" : "json parser implemented by go",`,
		`    "version": 0.1,`,
		`    "keyword": ["json", "parser", "go"]`,
		`}`,
	}
	parser := jp.NewParser(js)
	root, parseErr := parser.ParseValue()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}

	expected := make(map[string]interface{})
	expected["jsonparser"] = "json parser implemented by go"
	expected["version"] = 0.1
	expected["keyword"] = []interface{}{"json", "parser", "go"}

	if !reflect.DeepEqual(root.Evaluate(), expected) {
		tester.Error("failed simple json test")
	}
}

func TestParseError(tester *testing.T) {
	js := []string{
		`{`,
		`    "jsonparser : "json parser implemented by go",`,
		`}`,
	}
	parser := jp.NewParser(js)
	_, parseErr := parser.ParseValue()
	if parseErr != nil {
		errMsg := parseErr.Error()
		if !(strings.Contains(errMsg, "line 1") && strings.Contains(errMsg, "col 19") &&
			strings.Contains(errMsg, `unexpected token "j"`) &&
			strings.Contains(errMsg, `but expected COLON(:)`)) {
			tester.Error("bad error message")
		}
	}

}
