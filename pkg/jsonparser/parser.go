package jsonparser

import (
	"strconv"
	"strings"

	"github.com/hayas1/jsonparser/pkg/jsonparser/ast"
)

type Parser struct {
	lexer *Lexer
}

func NewParser(jsonLines []string) Parser {
	runes := make([][]rune, len(jsonLines))
	for i, s := range jsonLines {
		runes[i] = []rune(s)
	}
	return Parser{&Lexer{runes, 0, 0}}
}

func (p *Parser) ParseObject() (ast.ObjectNode, error) {
	if _, err := p.lexer.Lex1RuneToken(LEFTBRACE); err != nil {
		return ast.ObjectNode{}, err
	}

	valueObject := make(map[ast.StringNode]ast.ValueNode, 0)
	for !p.lexer.IsSkipWsToken(RIGHTBRACE) {
		stringNode, err1 := p.ParseString()
		if err, ok := err1.(*UnexpectedTokenError); err1 != nil && ok {
			err.AddExpected(RIGHTBRACE)
			return ast.ObjectNode{}, err
		} else if err1 != nil {
			return ast.ObjectNode{}, err1
		}

		_, err2 := p.lexer.Lex1RuneToken(COLON)
		if err2 != nil {
			return ast.ObjectNode{}, err2
		}

		valueNode, err3 := p.ParseValue()
		if err3 != nil {
			return ast.ObjectNode{}, err3
		}

		// should be refactored?
		if !p.lexer.IsSkipWsToken(RIGHTBRACE) {
			_, err4 := p.lexer.Lex1RuneToken(COMMA)
			if err, ok := err4.(*UnexpectedTokenError); ok && !p.lexer.IsSkipWsToken(RIGHTBRACE) {
				err.AddExpected(RIGHTBRACE)
				return ast.ObjectNode{}, err
			}
		}

		valueObject[stringNode] = valueNode
	}

	if _, err := p.lexer.Lex1RuneToken(RIGHTBRACE); err != nil {
		return ast.ObjectNode{}, err
	}

	return ast.ObjectNode{ValueObject: valueObject}, nil
}

func (p *Parser) ParseArray() (ast.ArrayNode, error) {
	if _, err := p.lexer.Lex1RuneToken(LEFTBRACKET); err != nil {
		return ast.ArrayNode{}, err
	}

	valueArray := make([]ast.ValueNode, 0)
	for !p.lexer.IsSkipWsToken(RIGHTBRACKET) {
		valueNode, err1 := p.ParseValue()
		if err1 != nil {
			return ast.ArrayNode{}, err1
		}

		// should be refactored?
		if !p.lexer.IsSkipWsToken(RIGHTBRACKET) {
			_, err2 := p.lexer.Lex1RuneToken(COMMA)
			if err, ok := err2.(*UnexpectedTokenError); ok && !p.lexer.IsSkipWsToken(RIGHTBRACKET) {
				err.AddExpected(RIGHTBRACKET)
				return ast.ArrayNode{}, err
			}
		}

		valueArray = append(valueArray, valueNode)
	}

	if _, err := p.lexer.Lex1RuneToken(RIGHTBRACKET); err != nil {
		return ast.ArrayNode{}, err
	}

	return ast.ArrayNode{ValueArray: valueArray}, nil
}

func (p *Parser) ParseValue() (ast.ValueNode, error) {
	// TODO
	return ast.ValueNode{}, nil
}

func (p *Parser) ParseImmediate() (ast.ImmediateNode, error) {
	// TODO
	return ast.ImmediateNode{}, nil
}

func (p *Parser) ParseString() (ast.StringNode, error) {
	startRow, startCol, _ := p.lexer.CurrentCursor()
	if _, err := p.lexer.Lex1RuneToken(QUOTATION); err != nil {
		return ast.StringNode{}, err
	}

	var builder strings.Builder
	for c, eof := p.lexer.CurrentRune(); !p.lexer.IsCurrentToken(QUOTATION); c, eof = p.lexer.Next() {
		if eof != nil {
			return ast.StringNode{}, &UnexpectedEofError{p.lexer.row, p.lexer.col, "parse string"}
		} else if startRow != p.lexer.row {
			return ast.StringNode{}, &OpenStringLiteralError{startRow, startCol}
		}

		switch c {
		case '\\':
			if err := p.ParseStringEscapeSequence(&builder); err != nil {
				return ast.StringNode{}, err
			}
		default:
			builder.WriteRune(c)
		}
	}

	if _, err := p.lexer.Lex1RuneToken(QUOTATION); err != nil {
		return ast.StringNode{}, err
	}

	return ast.StringNode{Str: builder.String()}, nil
}

func (p *Parser) ParseStringEscapeSequence(builder *strings.Builder) error {
	cc, eof := p.lexer.Next()
	if eof != nil {
		return &UnexpectedEofError{p.lexer.row, p.lexer.col, "parse escape sequence"}
	}
	switch cc {
	case '"', '/', '\\':
		builder.WriteRune(cc)
	case 'b', 'f', 'n', 'r', 't':
		escaped, err := strconv.Unquote(`"\` + string(cc) + `"`)
		if err != nil {
			return err
		}
		builder.WriteString(escaped)
	case 'u':
		hexUc, err := p.lexer.LexU4hexDigits()
		if err != nil {
			return err
		}
		strUc, err := strconv.Unquote(`"\u` + hexUc + `"`)
		if err != nil {
			return err
		}
		builder.WriteString(strUc)
	default:
		return &UnknownEscapeSequenceError{p.lexer.row, p.lexer.col, cc}
	}
	return nil
}

func (p *Parser) ParseNumber() (ast.NumberNode, error) {
	// TODO
	return ast.NumberNode{}, nil
}
