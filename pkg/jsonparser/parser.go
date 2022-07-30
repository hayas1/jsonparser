package jsonparser

import (
	ast "github.com/hayas1/jsonparser/pkg/jsonparser/ast"
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
	for !p.lexer.IsObjectEnd() {
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
		if !p.lexer.IsObjectEnd() {
			_, err4 := p.lexer.Lex1RuneToken(COMMA)
			if err, ok := err4.(*UnexpectedTokenError); !p.lexer.IsObjectEnd() && ok {
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

func (p *Parser) ParseString() (ast.StringNode, error) {
	// TODO
	return ast.StringNode{}, nil
}

func (p *Parser) ParseValue() (ast.ValueNode, error) {
	// TODO
	return ast.ValueNode{}, nil
}
