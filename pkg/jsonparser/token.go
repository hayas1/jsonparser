package jsonparser

import "fmt"

type TokenType string

const (
	LEFTBRACE    TokenType = "LEFTBRACE({)"
	RIGHTBRACE   TokenType = "RIGHTBRACE(})"
	WHITESPACE   TokenType = "WHITESPACE( )"
	COLON        TokenType = "COLON(:)"
	COMMA        TokenType = "COMMA(,)"
	LEFTBRACKET  TokenType = "LEFTBRACKET(])"
	RIGHTBRACKET TokenType = "RIGHTBRACKET(])"

	UNKNOWN TokenType = "UNKNOWN"
)

type Token struct {
	Element   string
	TokenType TokenType
}

func (t Token) String() string {
	return fmt.Sprintf("<\"%s\": %s>", t.Element, t.TokenType)
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
	default:
		return Token{s, UNKNOWN}
	}
}
