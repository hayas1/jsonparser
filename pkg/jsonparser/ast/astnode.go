package ast

type AstNode interface {
	Evaluate() interface{}
	Dump(int) string
}
