package jsonparser

import "io"

type Lexer struct {
	js  [][]rune
	row int
	col int
}

func (l *Lexer) CurrentCursor() (int, int, error) {
	if l.row < len(l.js) {
		return l.row, l.col, nil
	} else {
		return len(l.js), 0, io.EOF
	}
}

func (l *Lexer) NextCursor(row int, col int) (int, int, error) {
	if row >= len(l.js) {
		return len(l.js), 0, io.EOF
	} else if col+1 < len(l.js[row]) {
		return row, col + 1, nil
	} else if row+1 < len(l.js) {
		return row + 1, 0, nil
	} else {
		return len(l.js), 0, io.EOF
	}
}

func (l *Lexer) GetSkipWsCursor() (int, int, error) {
	row, col, eof := l.CurrentCursor()
	for ; l.row < len(l.js) && IsWhitespace(l.js[row][col]); row, col, eof = l.NextCursor(row, col) {
		if eof != nil {
			return len(l.js), 0, eof
		}
	}
	return row, col, eof
}

func (l *Lexer) PeekSkipWsToken() Token {
	row, col, eof := l.GetSkipWsCursor()
	if eof == nil {
		return Tokenize(l.js[row][col])
	} else {
		return Token{"", EOF}
	}
}
func (l *Lexer) IsSkipWsToken(t TokenType) bool {
	return l.PeekSkipWsToken().TokenType == t
}

func (l *Lexer) Lex1RuneToken(expectedToken TokenType) (Token, error) {
	row, col, eof := l.GetSkipWsCursor()
	var token Token
	if eof == nil {
		token = Tokenize(l.js[row][col])
	} else {
		token = Token{"", EOF}
	}

	if token.TokenType == expectedToken {
		l.row, l.col, _ = l.NextCursor(row, col)
		return token, nil
	}
	return token, &UnexpectedTokenError{row, col, token, []TokenType{expectedToken}}
}
