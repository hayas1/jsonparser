package ast

import "strconv"

type NumberNode struct {
	Number string
}

func (n *NumberNode) Children() []*AstNode {
	return make([]*AstNode, 0)
}

func (n *NumberNode) Evaluate() interface{} {
	num := n.Float()
	return num
}

func (n *NumberNode) Integer() int64 {
	// TODO no via float
	num := n.Float()
	return int64(num)
}

func (n *NumberNode) Float() float64 {
	num, _ := strconv.ParseFloat(n.Number, 64)
	return num
}
