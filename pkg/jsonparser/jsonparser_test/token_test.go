package jsonparser_test

import (
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func testTokenize(tester *testing.T, tokenRune []rune, expectedTokenTypes []jp.TokenType, errorMsg string) {
	n := len(tokenRune)
	if n != len(expectedTokenTypes) {
		tester.Error("tokenStrings and expectedTokenTypes should have same length")
		return
	}
	isAllOk := true
	for i := 0; i < n; i++ {
		isAllOk = isAllOk && jp.Tokenize(tokenRune[i]).TokenType == expectedTokenTypes[i]
	}
	if !isAllOk {
		tester.Error(errorMsg)
	}
}

func TestTokenizeBrace(tester *testing.T) {
	leftbrace, rightbrace := '{', '}'
	braces := []rune{leftbrace, rightbrace}
	expected := []jp.TokenType{jp.LEFTBRACE, jp.RIGHTBRACE}
	testTokenize(tester, braces, expected, "\"{\", \"}\" should be brace token")
}

func TestTokenizeWhitespace(tester *testing.T) {
	space, linefeed, carriageReturn, tab := ' ', '\n', '\r', '\t'
	ws := []rune{space, linefeed, carriageReturn, tab}
	expected := []jp.TokenType{jp.WHITESPACE, jp.WHITESPACE, jp.WHITESPACE, jp.WHITESPACE}
	testTokenize(tester, ws, expected, "space, linefeed, carriage return, tab should be white space token")
}

func TestTokenizeColon(tester *testing.T) {
	colon := ':'
	colons := []rune{colon}
	expected := []jp.TokenType{jp.COLON}
	testTokenize(tester, colons, expected, "\":\" should be colon token")
}

func TestTokenizeComma(tester *testing.T) {
	comma := ','
	commas := []rune{comma}
	expected := []jp.TokenType{jp.COMMA}
	testTokenize(tester, commas, expected, "\",\" should be comma token")
}

func TestTokenizeBracket(tester *testing.T) {
	leftbracket, rightbracket := '[', ']'
	brackets := []rune{leftbracket, rightbracket}
	expected := []jp.TokenType{jp.LEFTBRACKET, jp.RIGHTBRACKET}
	testTokenize(tester, brackets, expected, "\"[\", \"]\" should be bracket token")
}

func TestTokenizeImmediate(tester *testing.T) {
	tru, fal, null := "true", "false", "null"
	immediate := []string{tru, fal, null}
	expected := []jp.TokenType{jp.TRUE, jp.FALSE, jp.NULL}
	isAllOk := true
	for i := 0; i < 3; i++ {
		isAllOk = isAllOk && jp.TokenizeImmediate(immediate[i]).TokenType == expected[i]
	}
	if !isAllOk {
		tester.Error("\"true\", \"false\", \"\" should be immediate token")
	}
}

func TestTokenizeQuotation(tester *testing.T) {
	quotation := '"'
	quotations := []rune{quotation}
	expected := []jp.TokenType{jp.QUOTATION}
	testTokenize(tester, quotations, expected, "\"\"\" should be quotation token")
}

func TestTokenizeString(tester *testing.T) {
	quo, rsl, sl, bs, ff, lf, cr, ht, uni := "\"", "\\", "/", "b", "f", "n", "r", "t", "u"
	stringParts := []string{quo, rsl, sl, bs, ff, lf, cr, ht, uni}
	expected := []jp.TokenType{jp.QUOTATION, jp.REVERSESOLIDUS, jp.SOLIDUS, jp.BACKSPACE, jp.FORMFEED, jp.LINEFEED,
		jp.CARRIAGERETURN, jp.HORIZONTALTAB, jp.UNICODE}
	for i := 0; i < len(stringParts); i++ {
		if jp.TokenizeString(stringParts[i]).TokenType != expected[i] {
			tester.Error(stringParts[i] + " should be " + string(expected[i]) + " token")
		}
	}

	if jp.TokenizeString("aaa").TokenType != jp.SUBSTRING {
		tester.Error("\"aaa\" should be substring type token")
	}
}

func TestTokenizeNumber(tester *testing.T) {
	digit, minus, plus, dot, exponent := '7', '-', '+', '.', 'e'
	numberParts := []rune{digit, minus, plus, dot, exponent}
	expected := []jp.TokenType{jp.DIGIT, jp.MINUS, jp.PLUS, jp.DOT, jp.EXPONENT}
	for i := 0; i < len(numberParts); i++ {
		if jp.TokenizeNumber(numberParts[i]).TokenType != expected[i] {
			tester.Error(string(numberParts[i]) + " should be " + string(expected[i]) + " token")
		}
	}

}
