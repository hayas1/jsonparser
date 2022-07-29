package ast

type StringNode struct {
	str string
}

func (n *StringNode) Children() []*AstNode {
	return make([]*AstNode, 0)
}

func (n *StringNode) Evaluate() interface{} {
	return n.String()
}

func (n *StringNode) String() interface{} {
	return n.str
}