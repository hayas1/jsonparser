package ast

type ValueNode struct {
	Child AstNode
}

func (n *ValueNode) Evaluate() interface{} {
	return n.Child.Evaluate()
}

func (n *ValueNode) Dump(indent int) string {
	return n.Child.Dump(indent)
}
