package jsonparser_test

import (
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func TestParseObject(tester *testing.T) {
	emptyObject := []string{"  \n{ \n \t \t } \t\t"}
	parser := jp.NewParser(emptyObject)
	objectNode, _ := parser.ParseObject()
	emptyMap, ok := objectNode.Evaluate().(map[string]interface{})
	if !ok || len(emptyMap) > 0 {
		tester.Error("emptyObject should be parsed as empty map")
	}
}
