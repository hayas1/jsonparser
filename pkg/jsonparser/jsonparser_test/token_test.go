package jsonparser_test

import (
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func testTokenize(tester *testing.T, tokenStrings []string, expectedTokenTypes []jp.TokenType, errorMsg string) {
	n := len(tokenStrings)
	if n != len(expectedTokenTypes) {
		tester.Error("tokenStrings and expectedTokenTypes should have same length")
		return
	}
	isAllOk := true
	for i := 0; i < n; i++ {
		isAllOk = isAllOk && jp.Tokenize(tokenStrings[i]).TokenType == expectedTokenTypes[i]
	}
	if !isAllOk {
		tester.Error(errorMsg)
	}
}

func TestTokenizeBrace(tester *testing.T) {
	leftbrace, rightbrace := "{", "}"
	braces := []string{leftbrace, rightbrace}
	expected := []jp.TokenType{jp.LEFTBRACE, jp.RIGHTBRACE}
	testTokenize(tester, braces, expected, "\"{\", \"}\" should be brace token")
}

func TestTokenizeWhitespace(tester *testing.T) {
	space, linefeed, carriageReturn, tab := " ", "\n", "\r\n", "\t"
	ws := []string{space, linefeed, carriageReturn, tab}
	expected := []jp.TokenType{jp.WHITESPACE, jp.WHITESPACE, jp.WHITESPACE, jp.WHITESPACE}
	testTokenize(tester, ws, expected, "space, linefeed, carriage return, tab should be white space token")
}

func TestTokenizeColon(tester *testing.T) {
	colon := ":"
	colons := []string{colon}
	expected := []jp.TokenType{jp.COLON}
	testTokenize(tester, colons, expected, "\":\" should be colon token")
}

func TestTokenizeComma(tester *testing.T) {
	comma := ","
	commas := []string{comma}
	expected := []jp.TokenType{jp.COMMA}
	testTokenize(tester, commas, expected, "\",\" should be comma token")
}

func TestTokenizeBracket(tester *testing.T) {
	leftbracket, rightbracket := "[", "]"
	brackets := []string{leftbracket, rightbracket}
	expected := []jp.TokenType{jp.LEFTBRACKET, jp.RIGHTBRACKET}
	testTokenize(tester, brackets, expected, "\"[\", \"]\" should be bracket token")
}

func TestTokenizeImmediate(tester *testing.T) {
	tru, fal, null := "true", "false", "null"
	immediate := []string{tru, fal, null}
	expected := []jp.TokenType{jp.TRUE, jp.FALSE, jp.NULL}
	testTokenize(tester, immediate, expected, "\"true\", \"false\", \"\" should be immediate token")
}

func TestTokenizeQuotation(tester *testing.T) {
	quotation := "\""
	quotations := []string{quotation}
	expected := []jp.TokenType{jp.QUOTATION}
	testTokenize(tester, quotations, expected, "\"\"\" should be quotation token")
}

func TestTokenizeString(tester *testing.T) {
	str1, str2 := "string", "\"string\""
	strings := []string{str1, str2}
	expected := []jp.TokenType{jp.STRING, jp.STRING}
	for i := 0; i < len(strings); i++ {
		if jp.TokenizeString(strings[i]).TokenType != expected[i] {
			tester.Error("TokenizeString function must return STRING type token")
		}
		if jp.TokenizeString(strings[i]).Element != strings[i] {
			tester.Error("STRING type token's Element is origin string")
		}
	}
}

func TestTokenizeNumber(tester *testing.T) {
	num100 := jp.TokenizeNumber("100")
	if num100.TokenType != jp.NUMBER {
		tester.Error("TokenizeNumber function must return NUMBER type token")
	}
	if hundred, err := num100.Integer(); !(hundred == 100 && err == nil) {
		tester.Error("100 must be evaluated as integer 100")
	}

	num1p5 := jp.TokenizeNumber("1.5")
	if num1p5.TokenType != jp.NUMBER {
		tester.Error("TokenizeNumber function must return NUMBER type token")
	}
	if oneFive, err := num1p5.Float(); !(oneFive == 1.5 && err == nil) {
		tester.Error("1.5 must be evaluated as float 1.5")
	}

	num1e5 := jp.TokenizeNumber("1E5")
	if num1e5.TokenType != jp.NUMBER {
		tester.Error("TokenizeNumber function must return NUMBER type token")
	}
	if hundredThousand, err := num1e5.Integer(); !(hundredThousand == 100000 && err == nil) {
		tester.Error("1E5 must be evaluated as integer 100000")
	}

}
