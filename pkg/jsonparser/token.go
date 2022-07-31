package jsonparser

import (
	"fmt"
)

type TokenType string

const (
	LEFTBRACE    TokenType = "LEFTBRACE({)"
	RIGHTBRACE   TokenType = "RIGHTBRACE(})"
	WHITESPACE   TokenType = "WHITESPACE( )"
	COLON        TokenType = "COLON(:)"
	COMMA        TokenType = "COMMA(,)"
	LEFTBRACKET  TokenType = "LEFTBRACKET([)"
	RIGHTBRACKET TokenType = "RIGHTBRACKET(])"
	TRUE         TokenType = "TRUE(true)"
	FALSE        TokenType = "FALSE(false)"
	NULL         TokenType = "NULL(null)"

	QUOTATION      TokenType = "QUOTATION(\")"
	REVERSESOLIDUS TokenType = "REVERSESOLIDUS(\\)"
	SOLIDUS        TokenType = "SOLIDUS(/)"
	BACKSPACE      TokenType = "BACKSPACE(b)"
	FORMFEED       TokenType = "FORMFEED(f)"
	LINEFEED       TokenType = "LINEFEED(n)"
	CARRIAGERETURN TokenType = "CARRIAGERETURN(r)"
	HORIZONTALTAB  TokenType = "HORIZONTALTAB(t)"
	UNICODE        TokenType = "UNICODE(u)"
	SUBSTRING      TokenType = "SUBSTRING(\"...\")"

	DIGIT    TokenType = "DIGIT(0-9)"
	MINUS    TokenType = "MINUS(-)"
	PLUS     TokenType = "PLUS(+)"
	DOT      TokenType = "DOT(.)"
	EXPONENT TokenType = "EXPONENT(E)"
	DIGITS   TokenType = "DIGITS([0-9]+)"

	UNKNOWN TokenType = "UNKNOWN"
	EOF     TokenType = "EOF"
)

type Token struct {
	Element   string
	TokenType TokenType
}

func (t Token) String() string {
	return fmt.Sprintf("<\"%s\": %s>", t.Element, t.TokenType)
}

func IsWhitespace(c rune) bool {
	switch c {
	case ' ', '\n', '\r', '\t':
		return true
	default:
		return false
	}
}

func Tokenize(c rune) Token {
	switch c {
	case '{':
		return Token{string(c), LEFTBRACE}
	case '}':
		return Token{string(c), RIGHTBRACE}
	case ' ', '\n', '\r', '\t':
		return Token{string(c), WHITESPACE}
	case ':':
		return Token{string(c), COLON}
	case ',':
		return Token{string(c), COMMA}
	case '[':
		return Token{string(c), LEFTBRACKET}
	case ']':
		return Token{string(c), RIGHTBRACKET}
	case '"':
		return Token{string(c), QUOTATION}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return Token{string(c), DIGIT}
	case '-':
		return Token{string(c), MINUS}
	case '+':
		return Token{string(c), PLUS}
	case '.':
		return Token{string(c), DOT}
	case 'e', 'E':
		return Token{string(c), EXPONENT}
	default:
		return Token{string(c), UNKNOWN}
	}
}

func TokenizeImmediate(s string) Token {
	switch s {
	case "true":
		return Token{s, TRUE}
	case "false":
		return Token{s, FALSE}
	case "null":
		return Token{s, NULL}
	default:
		return Token{s, UNKNOWN}
	}
}

func TokenizeString(s string) Token {
	switch s {
	case "\"":
		return Token{s, QUOTATION}
	case "\\":
		return Token{s, REVERSESOLIDUS}
	case "/":
		return Token{s, SOLIDUS}
	case "b":
		return Token{s, BACKSPACE}
	case "f":
		return Token{s, FORMFEED}
	case "n":
		return Token{s, LINEFEED}
	case "r":
		return Token{s, CARRIAGERETURN}
	case "t":
		return Token{s, HORIZONTALTAB}
	case "u":
		return Token{s, UNICODE}
	default:
		return Token{s, SUBSTRING}
	}
}

func TokenizeNumber(c rune) Token {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return Token{string(c), DIGIT}
	case '-':
		return Token{string(c), MINUS}
	case '+':
		return Token{string(c), PLUS}
	case '.':
		return Token{string(c), DOT}
	case 'e', 'E':
		return Token{string(c), EXPONENT}
	default:
		return Token{string(c), UNKNOWN}
	}
}
