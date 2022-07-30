package ast

type ArrayNode struct {
	ValueArray []ValueNode
}

func (n *ArrayNode) Children() []AstNode {
	// TODO
	return make([]AstNode, 0)
}

func (n *ArrayNode) Evaluate() interface{} {
	return n.Array()
}

func (n *ArrayNode) Array() []interface{} {
	array := make([]interface{}, len(n.ValueArray))
	for i := 0; i < len(n.ValueArray); i++ {
		array[i] = n.ValueArray[i].Evaluate()
	}
	return array
}
