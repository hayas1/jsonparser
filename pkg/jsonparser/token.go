package jsonparser

import "fmt"

type TokenType string

const (
	LEFTBRACE  TokenType = "{"
	RIGHTBRACE TokenType = "}"
	WHITESPACE TokenType = " "
)

type Token struct {
	Element    string
	Token_type TokenType
}

func (t *Token) String() string {
	return fmt.Sprintf("[\"%s\": %s]", t.Element, t.Token_type)
}

func Tokenize(s string) Token {
	return Token{s, LEFTBRACE}
}
