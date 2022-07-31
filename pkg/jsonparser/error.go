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

type UnexpectedEofError struct {
	row  int
	col  int
	when string
}

func (e *UnexpectedEofError) Error() string {
	return fmt.Sprintf("line %d (col %d): unexpected EOF, during %s", e.row, e.col, e.when)
}

type UnexpectedLinefeedError struct {
	row  int
	col  int
	when string
}

func (e *UnexpectedLinefeedError) Error() string {
	return fmt.Sprintf("line %d (col %d): unexpected line feed, during %s", e.row, e.col, e.when)
}

type OpenStringLiteralError struct {
	row int
	col int
}

func (e *OpenStringLiteralError) Error() string {
	return fmt.Sprintf("line %d (col %d): string literal should be closed by \"\"\"", e.row, e.col)
}

type UnknownEscapeSequenceError struct {
	row            int
	col            int
	escapeSequence rune
}

func (e *UnknownEscapeSequenceError) Error() string {
	return fmt.Sprintf("line %d (col %d): unknown escape sequence %v", e.row, e.col, e.escapeSequence)
}

type CannotParseUnicodeError struct {
	row int
	col int
	msg string
}

func (e *CannotParseUnicodeError) Error() string {
	return fmt.Sprintf("line %d (col %d): unknown escape sequence %s", e.row, e.col, e.msg)
}

type UnknownImmediatePrefix struct {
	row    int
	col    int
	prefix string
}

func (e *UnknownImmediatePrefix) Error() string {
	return fmt.Sprintf("line %d (col %d): unknown immediate start with %s", e.row, e.col, e.prefix)
}
