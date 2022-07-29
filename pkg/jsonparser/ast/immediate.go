package ast

type ImmediateNode struct {
	// TODO enum? true or false or null
}

func (n *ImmediateNode) Children() []*AstNode {
	return make([]*AstNode, 0)
}

func (n *ImmediateNode) Evaluate() interface{} {
	// TODO
	return make([]interface{}, 0)
}

func (n *ImmediateNode) True() bool {
	// TODO
	return true
}

func (n *ImmediateNode) False() bool {
	// TODO
	return false
}

func (n *ImmediateNode) Null() struct{} {
	// TODO
	return struct{}{}
}
