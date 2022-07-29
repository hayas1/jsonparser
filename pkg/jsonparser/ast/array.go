package ast

type ArrayNode struct {
	Values []ValueNode
}

func (n *ArrayNode) Children() []AstNode {
	// TODO
	return make([]AstNode, 0)
}

func (n *ArrayNode) Evaluate() interface{} {
	// TODO
	return make([]interface{}, 0)
}

func (n *ArrayNode) Array() []interface{} {
	arr := make([]interface{}, len(n.Values))
	for i := 0; i < len(n.Values); i++ {
		arr[i] = n.Values[i].Evaluate()
	}
	return arr
}
