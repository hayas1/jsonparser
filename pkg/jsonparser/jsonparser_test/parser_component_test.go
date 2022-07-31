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
		tester.Error("emptyObject should be parsed as empty map, but length ", len(emptyMap))
	}
}

func TestParseEmptyArray(tester *testing.T) {
	emptyArray := []string{"  \t \r\n[ \n \t \t ] \t\t"}
	parser := jp.NewParser(emptyArray)
	arrayNode, parseErr := parser.ParseArray()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	emptyArr := arrayNode.Array()
	if len(emptyArr) > 0 {
		tester.Error("emptyObject should be parsed as empty map, but length ", len(emptyArr))
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
		tester.Error("\"json parser go\" should be parsed as \"json parser go\", but as", str)
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
		tester.Error("\"json parser \"go\"\" should be parsed as \"json parser \"go\"\", but as", str)
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
		tester.Error("\"json parser\t\"g\u00f3\"😶\" should be parsed as \"json parser\t\"gó\"😶\", but as", str)
	}
}

func TestParseInteger(tester *testing.T) {
	simpleInteger := []string{"123456789"}
	parser := jp.NewParser(simpleInteger)
	integerNode, parseErr := parser.ParseNumber()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	if integerNode.Integer() != 123456789 {
		tester.Error("\"123456789\" should be parsed as 123456789, but as ", integerNode.Integer())
	}
}

func TestParseMinusInteger(tester *testing.T) {
	simpleInteger := []string{"-123456789"}
	parser := jp.NewParser(simpleInteger)
	integerNode, parseErr := parser.ParseNumber()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	if integerNode.Integer() != -123456789 {
		tester.Error("\"-123456789\" should be parsed as -123456789, but as", integerNode.Integer())
	}
}

func TestParseFloat(tester *testing.T) {
	simpleFloat := []string{"1234.56789"}
	parser := jp.NewParser(simpleFloat)
	floatNode, parseErr := parser.ParseNumber()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	if floatNode.Float() != 1234.56789 {
		tester.Error("\"1234.56789\" should be parsed as 1234.56789, but as", floatNode.Float())
	}
}

func TestParseExponent(tester *testing.T) {
	simpleExponent := []string{"3e5"}
	parser := jp.NewParser(simpleExponent)
	exponentNode, parseErr := parser.ParseNumber()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	if exponentNode.Integer() != 300000 {
		tester.Error("\"3e5\" should be parsed as 300000, but as", exponentNode.Integer())
	}
}

func TestParseMinusExponent(tester *testing.T) {
	simpleExponent := []string{"5E-1"}
	parser := jp.NewParser(simpleExponent)
	exponentNode, parseErr := parser.ParseNumber()
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}
	if exponentNode.Float() != 0.5 {
		tester.Error("\"5E-1\" should be parsed as 0.5, but as", exponentNode.Float())
	}
}
