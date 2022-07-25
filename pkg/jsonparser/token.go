package jsonparser

import (
	"fmt"
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

	QUOTATION      TokenType = "QUOTATION(\")"
	REVERSESOLIDUS TokenType = "REVERSESOLIDUS(\\)"
	SOLIDUS        TokenType = "SOLIDUS(/)"
	BACKSPACE      TokenType = "BACKSPACE(b)"
	FORMFEED       TokenType = "FORMFEED(f)"
	LINEFEED       TokenType = "LINEFEED(n)"
	CARRIAGERETURN TokenType = "CARRIAGERETURN(r)"
	HORIZONTALTAB  TokenType = "HORIZONTALTAB(t)"
	UNICODE        TokenType = "UNICODE(u)"
	SUBSTRING      TokenType = "SUBSTRING(\"...\")"

	DIGIT    TokenType = "DIGIT(0-9)"
	MINUS    TokenType = "MINUS(-)"
	PLUS     TokenType = "PLUS(+)"
	DOT      TokenType = "DOT(.)"
	EXPONENT TokenType = "EXPONENT(E)"

	UNKNOWN TokenType = "UNKNOWN"
)

type Token struct {
	Element   string
	TokenType TokenType
}

func (t Token) String() string {
	return fmt.Sprintf("<\"%s\": %s>", t.Element, t.TokenType)
}

func IsWhitespace(c rune) bool {
	switch c {
	case ' ', '\n', '\r', '\t':
		return true
	default:
		return false
	}
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
	switch s {
	case "\"":
		return Token{s, QUOTATION}
	case "\\":
		return Token{s, REVERSESOLIDUS}
	case "/":
		return Token{s, SOLIDUS}
	case "b":
		return Token{s, BACKSPACE}
	case "f":
		return Token{s, FORMFEED}
	case "n":
		return Token{s, LINEFEED}
	case "r":
		return Token{s, CARRIAGERETURN}
	case "t":
		return Token{s, HORIZONTALTAB}
	case "u":
		return Token{s, UNICODE}
	default:
		return Token{s, SUBSTRING}
	}
}

func TokenizeNumber(s string) Token {
	switch s {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return Token{s, DIGIT}
	case "-":
		return Token{s, MINUS}
	case "+":
		return Token{s, PLUS}
	case ".":
		return Token{s, DOT}
	case "e", "E":
		return Token{s, EXPONENT}
	default:
		return Token{s, UNKNOWN}
	}
}
