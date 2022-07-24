package jsonparser

import (
	"fmt"
	"strconv"
)

type TokenType string

const (
	LEFTBRACE    TokenType = "LEFTBRACE({)"
	RIGHTBRACE   TokenType = "RIGHTBRACE(})"
	WHITESPACE   TokenType = "WHITESPACE( )"
	COLON        TokenType = "COLON(:)"
	COMMA        TokenType = "COMMA(,)"
	LEFTBRACKET  TokenType = "LEFTBRACKET(])"
	RIGHTBRACKET TokenType = "RIGHTBRACKET(])"
	TRUE         TokenType = "TRUE(true)"
	FALSE        TokenType = "FALSE(false)"
	NULL         TokenType = "NULL(null)"

	QUOTATION TokenType = "QUOTATION(\")"
	STRING    TokenType = "STRING"
	// REVERSESOLIDUS TokenType = "REVERSESOLIDUS(\\)"
	// SOLIDUS        TokenType = "SOLIDUS(/)"
	// BACKSPACE      TokenType = "BACKSPACE(b)"
	// FORMFEED       TokenType = "FORMFEED(f)"
	// LINEFEED       TokenType = "LINEFEED(n)"
	// CARRIAGERETURN TokenType = "CARRIAGERETURN(r)"
	// HORIZONTALTAB  TokenType = "HORIZONTALTAB(t)"
	// UNICODE        TokenType = "UNICODE(u)"

	NUMBER TokenType = "NUMBER"
	// MINUS    TokenType = "MINUS(-)"
	// PLUS     TokenType = "PLUS(+)"
	// DOT      TokenType = "DOT(.)"
	// EXPONENT TokenType = "EXPONENT(E)"

	UNKNOWN TokenType = "UNKNOWN"
)

type Token struct {
	Element   string
	TokenType TokenType
}

func (t Token) String() string {
	return fmt.Sprintf("<\"%s\": %s>", t.Element, t.TokenType)
}

func (t Token) Float() (float64, error) {
	return strconv.ParseFloat(t.Element, 64)
}

func (t Token) Integer() (int64, error) {
	num, err := t.Float()
	return int64(num), err
}

func Tokenize(s string) Token {
	switch s {
	case "{":
		return Token{s, LEFTBRACE}
	case "}":
		return Token{s, RIGHTBRACE}
	case " ", "\n", "\r\n", "\t":
		return Token{s, WHITESPACE}
	case ":":
		return Token{s, COLON}
	case ",":
		return Token{s, COMMA}
	case "[":
		return Token{s, LEFTBRACKET}
	case "]":
		return Token{s, RIGHTBRACKET}
	case "true":
		return Token{s, TRUE}
	case "false":
		return Token{s, FALSE}
	case "null":
		return Token{s, NULL}
	case "\"":
		return Token{s, QUOTATION}
	default:
		return Token{s, UNKNOWN}
	}
}

func TokenizeString(s string) Token {
	return Token{s, STRING}
}

func TokenizeNumber(s string) Token {
	return Token{s, NUMBER}
}
