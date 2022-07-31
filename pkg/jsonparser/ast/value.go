package ast

type ValueNode struct {
	Child AstNode
}

func (n *ValueNode) Children() []*AstNode {
	return []*AstNode{&n.Child}
}

func (n *ValueNode) Evaluate() interface{} {
	return n.Child.Evaluate()
}
