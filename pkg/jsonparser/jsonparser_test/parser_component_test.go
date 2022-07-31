package jsonparser_test

import (
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func TestParseEmptyObject(tester *testing.T) {
	emptyObject := []string{"{}"}
	parser := jp.NewParser(emptyObject)
	objectNode, parseErr := parser.ParseObject()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	emptyMap := objectNode.Object()
	if len(emptyMap) > 0 {
		tester.Error("emptyObject should be parsed as empty map")
	}
}

func TestParseEmptyArray(tester *testing.T) {
	emptyArray := []string{"  \t \r\n[ \n \t \t ] \t\t"}
	parser := jp.NewParser(emptyArray)
	arrayNode, parseErr := parser.ParseArray()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	emptyMap := arrayNode.Array()
	if len(emptyMap) > 0 {
		tester.Error("emptyObject should be parsed as empty map")
	}
}

func TestParseSimpleString(tester *testing.T) {
	simpleString := []string{`"json parser go"`}
	parser := jp.NewParser(simpleString)
	stringNode, parseErr := parser.ParseString()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	str := stringNode.String()
	if str != "json parser go" {
		tester.Error("\"json parser go\" should be parsed as \"json parser go\"")
	}
}

func TestParseSimpleEscapedString(tester *testing.T) {
	simpleString := []string{`"json parser \"go\""`}
	parser := jp.NewParser(simpleString)
	stringNode, parseErr := parser.ParseString()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	str := stringNode.String()
	if str != "json parser \"go\"" {
		tester.Error("\"json parser \"go\"\" should be parsed as \"json parser \"go\"\"")
	}
}

func TestParseEscapedString(tester *testing.T) {
	simpleString := []string{`"json parser\t\"g\u00f3\"😶"`}
	parser := jp.NewParser(simpleString)
	stringNode, parseErr := parser.ParseString()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	str := stringNode.String()
	if str != "json parser\t\"g\u00f3\"😶" {
		tester.Error("\"json parser\t\"g\u00f3\"😶\" should be parsed as \"json parser\t\"gó\"😶\"")
	}
}
