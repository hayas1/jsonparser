package ast

type ValueNode struct {
	// TODO enum?
}

func (n *ValueNode) Children() []*AstNode {
	// TODO
	return make([]*AstNode, 0)
}

func (n *ValueNode) Evaluate() interface{} {
	// TODO
	return make([]interface{}, 0)
}

// TODO
// func (n *ValueNode) Integer() int {}
// func (n *ValueNode) Object() map[string]interface{} {}
// AND SO ON...
