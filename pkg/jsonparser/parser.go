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
	if _, err := p.lexer.LexObjectHead(); err != nil {
		return ast.ObjectNode{}, err
	}
	objectMap := make(map[ast.StringNode]ast.ValueNode, 0)
	for p.lexer.IsNextString() {
		stringNode, err1 := p.ParseString()
		if err1 != nil {
			return ast.ObjectNode{}, err1
		}
		_, err2 := p.lexer.LexObjectSplit()
		if err2 != nil {
			return ast.ObjectNode{}, err2
		}
		objectNode, err3 := p.ParseValue()
		if err3 != nil {
			return ast.ObjectNode{}, err3
		}
		objectMap[stringNode] = objectNode
	}

	return ast.ObjectNode{}, nil
}

func (p *Parser) ParseString() (ast.StringNode, error) {
	// TODO
	return ast.StringNode{}, nil
}

func (p *Parser) ParseValue() (ast.ValueNode, error) {
	// TODO
	return ast.ValueNode{}, nil
}
