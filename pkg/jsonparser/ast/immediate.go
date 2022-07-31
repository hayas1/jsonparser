package ast

import "errors"

type ImmediateNode struct {
	Immediate string
}

func (n *ImmediateNode) Children() []*AstNode {
	return make([]*AstNode, 0)
}

func (n *ImmediateNode) Evaluate() interface{} {
	switch n.Immediate {
	case "true":
		return n.True()
	case "false":
		return n.False()
	case "null":
		return n.Null()
	default:
		return errors.New("unexpected Immediate: " + n.Immediate)
	}
}

func (n *ImmediateNode) True() bool {
	return n.Immediate == "true"
}

func (n *ImmediateNode) False() bool {
	return n.Immediate == "false"
}

func (n *ImmediateNode) Null() bool {
	// TODO use enum constant and better interface
	return n.Immediate == "null"
}
