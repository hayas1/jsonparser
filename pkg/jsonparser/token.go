package jsonparser

import "fmt"

type TokenType string

const (
	LEFTBRACE  TokenType = "{"
	RIGHTBRACE TokenType = "}"
	WHITESPACE TokenType = " "
)

type Token struct {
	element    string
	token_type TokenType
}

func (t *Token) String() string {
	return fmt.Sprintf("[\"%s\": %s]", t.element, t.token_type)
}

func Tokenize(s string) Token {
	return Token{s, LEFTBRACE}
}
