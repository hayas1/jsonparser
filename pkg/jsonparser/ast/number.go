package ast

type NumberNode struct {
	// TODO enum? int or float
}

func (n *NumberNode) Children() []*AstNode {
	return make([]*AstNode, 0)
}

func (n *NumberNode) Evaluate() interface{} {
	// TODO
	return make([]interface{}, 0)
}

func (n *NumberNode) Integer() int {
	// TODO
	return 0
}

func (n *NumberNode) Float() float64 {
	// TODO
	return 0
}
