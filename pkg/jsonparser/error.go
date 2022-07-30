package jsonparser

import (
	"fmt"
	"strings"
)

type UnexpectedTokenError struct {
	row      int
	col      int
	token    Token
	expected []TokenType
}

func (e *UnexpectedTokenError) Error() string {
	expectList := make([]string, len(e.expected))
	for i := 0; i < len(e.expected); i++ {
		expectList[i] = string(e.expected[i])
	}
	expected := fmt.Sprint("", strings.Join(expectList, " or "))
	return fmt.Sprintf("line %d (col %d): unexpected token \"%s\", but expected %s", e.row, e.col, e.token.Element, expected)
}

func (e *UnexpectedTokenError) AddExpected(expected ...TokenType) {
	e.expected = append(e.expected, expected...)
}
