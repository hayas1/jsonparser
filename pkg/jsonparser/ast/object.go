package ast

type ObjectNode struct {
	ValueObject map[StringNode]ValueNode
}

func (n *ObjectNode) Children() []*AstNode {
	// TODO
	return make([]*AstNode, 0)
}

func (n *ObjectNode) Evaluate() interface{} {
	return n.Object()
}

func (n *ObjectNode) Object() map[string]interface{} {
	object := make(map[string]interface{})
	for s, v := range n.ValueObject {
		object[s.Evaluate().(string)] = v.Evaluate()
	}
	return object
}
