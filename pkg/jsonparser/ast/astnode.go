package ast

type AstNode interface {
	Children() []*AstNode
	Evaluate() interface{}
}
