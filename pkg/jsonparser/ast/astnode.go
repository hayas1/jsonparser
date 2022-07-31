package ast

type AstNode interface {
	Evaluate() interface{}
}
