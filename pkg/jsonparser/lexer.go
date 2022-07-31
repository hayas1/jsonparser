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
func (l *Lexer) CurrentRune() (rune, error) {
	if l.IsCursorEOF(l.row, l.col) {
		return '\x00', io.EOF
	} else {
		return l.js[l.row][l.col], nil
	}
}
func (l *Lexer) IsCurrentToken(t TokenType) bool {
	c, eof := l.CurrentRune()
	var token Token
	if eof == nil {
		token = Tokenize(c)
	} else {
		token = Token{"", EOF}
	}
	return token.TokenType == t
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
func (l *Lexer) Next() (rune, error) {
	row, col, eof := l.CurrentCursor()
	if eof == nil {
		l.row, l.col, eof = l.NextCursor(row, col)
		// return l.js[row][col], nil
		if eof == nil {
			return l.js[l.row][l.col], nil
		} else {
			return '\x00', io.EOF
		}
	} else {
		return '\x00', io.EOF
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

func (l *Lexer) LexU4hexDigits() (string, error) {
	row, col, eof := l.CurrentCursor()
	if eof != nil {
		return "", &UnexpectedEofError{row, col, "parse unicode"}
	} else if len(l.js[row]) < col+5 {
		return "", &UnexpectedLinefeedError{row, col, "parse unicode"}
	}

	hex4 := l.js[row][col+1 : col+5]
	for _, uc := range hex4 {
		if !('0' <= uc && uc <= '9' || 'a' <= uc && uc <= 'f' || 'A' <= uc && uc <= 'F') {
			return "", &CannotParseUnicodeError{row, col, "unicode must be 4 hex digits(0-9,a-f)"}
		}
	}

	l.row, l.col = row, col+4
	return string(hex4), nil
}

func (l *Lexer) IsCurrentNumberToken(t TokenType) bool {
	c, eof := l.CurrentRune()
	if eof == nil {
		return TokenizeNumber(c).TokenType == t
	} else {
		return false
	}
}

func (l *Lexer) LexImmediate() (Token, error) {
	row, col, eof := l.CurrentCursor()
	if eof != nil {
		return Token{"", EOF}, eof
	}
	switch l.js[row][col] {
	case 't':
		if len(l.js[row]) < col+4 {
			return Token{"", UNKNOWN}, &UnexpectedLinefeedError{row, col, "parse immediate"}
		} else if tru := string(l.js[row][col : col+4]); tru != "true" {
			t := Token{tru, UNKNOWN}
			return t, &UnexpectedTokenError{row, col, t, []TokenType{TRUE}}
		} else {
			l.row, l.col = row, col+4
			return TokenizeImmediate(tru), nil
		}
	case 'f':
		if len(l.js[row]) < col+5 {
			return Token{"", UNKNOWN}, &UnexpectedLinefeedError{row, col, "parse immediate"}
		} else if fal := string(l.js[row][col : col+5]); fal != "false" {
			t := Token{fal, UNKNOWN}
			return t, &UnexpectedTokenError{row, col, t, []TokenType{TRUE}}
		} else {
			l.row, l.col = row, col+5
			return TokenizeImmediate(fal), nil
		}
	case 'n':
		if len(l.js[row]) < col+4 {
			return Token{"", UNKNOWN}, &UnexpectedLinefeedError{row, col, "parse immediate"}
		} else if null := string(l.js[row][col : col+4]); null != "null" {
			t := Token{null, UNKNOWN}
			return t, &UnexpectedTokenError{row, col, t, []TokenType{TRUE}}
		} else {
			l.row, l.col = row, col+4
			return TokenizeImmediate(null), nil
		}
	default:
		t := string(l.js[row][col])
		return Token{t, UNKNOWN}, &UnknownImmediatePrefix{row, col, t}
	}
}
