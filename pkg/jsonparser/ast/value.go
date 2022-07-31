package ast

type ValueNode struct {
	Child AstNode
}

func (n *ValueNode) Evaluate() interface{} {
	return n.Child.Evaluate()
}
