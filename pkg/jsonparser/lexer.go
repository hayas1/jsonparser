package jsonparser

import "io"

type Lexer struct {
	js  [][]rune
	row int
	col int
}

func (l *Lexer) CurrentCursor() (int, int) {
	return l.row, l.col
}
func (l *Lexer) Current() rune {
	return l.js[l.row][l.col]
}

func (l *Lexer) NextCursor(row int, col int) (int, int, error) {
	if col+1 < len(l.js[row]) {
		return row, col + 1, nil
	} else if row+1 < len(l.js) {
		return row + 1, 0, nil
	} else {
		return len(l.js), 0, io.EOF
	}
}
func (l *Lexer) GetNextCursor() (int, int, error) {
	return l.NextCursor(l.row, l.col)
}
func (l *Lexer) Peek() (rune, error) {
	row, col, eof := l.GetNextCursor()
	return l.js[row][col], eof
}

func (l *Lexer) GetNextNotWsCursor() (int, int, error) {
	row, col, eof := l.GetNextCursor()
	for ; IsWhitespace(l.js[row][col]); row, col, eof = l.NextCursor(row, col) {
		if eof != nil {
			return 0, 0, eof
		}
	}
	return row, col, nil
}
func (l *Lexer) skipWhitespace() error {
	row, col, eof := l.GetNextNotWsCursor()
	if eof != nil {
		l.row, l.col = row, col
		return eof
	} else {
		l.row, l.col = 0, 0
		return nil
	}
}
func (l *Lexer) PeekNextNotWsToken() Token {
	row, col, eof := l.GetNextCursor()
	for ; IsWhitespace(l.js[row][col]); row, col, eof = l.NextCursor(row, col) {
		if eof != nil {
			return Token{"", EOF}
		}
	}
	return Tokenize(l.js[row][col])
}
func (l *Lexer) IsObjectEnd() bool {
	return l.PeekNextNotWsToken().TokenType == LEFTBRACE
}

func (l *Lexer) Lex1RuneToken(expectedToken TokenType) (Token, error) {
	if eof := l.skipWhitespace(); eof != nil {
		return Token{"", EOF}, eof
	}
	token := Tokenize(l.Current())
	if token.TokenType == expectedToken {
		return token, nil
	} else {
		row, col := l.CurrentCursor()
		return token, &UnexpectedTokenError{row, col, token, []TokenType{expectedToken}}
	}
}

// func (l *Lexer) LexObjectStart() (Token, error) {
// 	return l.Lex1RuneToken(LEFTBRACE)
// }
// func (l *Lexer) LexObjectMap() (Token, error) {
// 	return l.Lex1RuneToken(COLON)
// }
// func (l *Lexer) LexObjectSplit() (Token, error) {
// 	return l.Lex1RuneToken(COMMA)
// }
// func (l *Lexer) LexObjectEnd() (Token, error) {
// 	return l.Lex1RuneToken(RIGHTBRACE)
// }
