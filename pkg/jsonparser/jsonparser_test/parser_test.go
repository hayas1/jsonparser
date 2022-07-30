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
