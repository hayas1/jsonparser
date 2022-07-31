package ast

import "strconv"

type StringNode struct {
	Str string
}

func (n *StringNode) Evaluate() interface{} {
	return n.String()
}

func (n *StringNode) String() string {
	return n.Str
}

func (n *StringNode) Dump(indent int) string {
	return strconv.Quote(n.Str)
}
