package jsonparser_test

import (
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func TestTokenizeBrace(tester *testing.T) {
	leftbrace, rightbrace := "{", "}"
	tl, tr := jp.Tokenize(leftbrace), jp.Tokenize(rightbrace)
	if !(tl.TokenType == jp.LEFTBRACE && tr.TokenType == jp.RIGHTBRACE) {
		tester.Error("\"{\" and \"}\" should be brace token")
	}
}

func TestTokenizeWhitespace(tester *testing.T) {
	space, linefeed, carriageReturn, tab := " ", "\n", "\r\n", "\t"
	ts, tl, tc, tt := jp.Tokenize(space), jp.Tokenize(linefeed), jp.Tokenize(carriageReturn), jp.Tokenize(tab)
	isAllWhiteSpace := true
	for _, t := range []jp.Token{ts, tl, tc, tt} {
		isAllWhiteSpace = isAllWhiteSpace && t.TokenType == jp.WHITESPACE
	}
	if !isAllWhiteSpace {
		tester.Error("space, linefeed, carriage return, tab should be white space token")
	}
}

func TestTokenizeColon(tester *testing.T) {
	colon := ":"
	tc := jp.Tokenize(colon)
	if !(tc.TokenType == jp.COLON) {
		tester.Error("\":\" should be brace token")
	}
}
