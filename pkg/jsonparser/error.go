package jsonparser

import "fmt"

type UnexpectedTokenError struct {
	row   int
	col   int
	token Token
}

func (e *UnexpectedTokenError) Error() string {
	return fmt.Sprintf("line %d(col %d): unexpected token %s", e.row, e.col, e.token.Element)
}
