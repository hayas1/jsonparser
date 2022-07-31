package jsonparser

import "io"

type Lexer struct {
	js  [][]rune
	row int // if len(js) <= row, reach EOF
	col int
}

func (l *Lexer) IsCursorEOF(row int, col int) bool {
	if row < len(l.js) {
		return false
	} else {
		return true
	}
}
func (l *Lexer) CurrentCursor() (int, int, error) {
	if l.IsCursorEOF(l.row, l.col) {
		return len(l.js), 0, io.EOF
	} else {
		return l.row, l.col, nil
	}
}
func (l *Lexer) NextCursor(row int, col int) (int, int, error) {
	if l.IsCursorEOF(l.row, l.col) {
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
	for ; !l.IsCursorEOF(row, col) && IsWhitespace(l.js[row][col]); row, col, eof = l.NextCursor(row, col) {
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
