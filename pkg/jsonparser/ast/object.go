package ast

type ObjectNode struct {
	objectMap map[StringNode]ValueNode
}

func (n *ObjectNode) Children() []*AstNode {
	// TODO
	return make([]*AstNode, 0)
}

func (n *ObjectNode) Evaluate() interface{} {
	return n.Object()
}

func (n *ObjectNode) Object() map[string]interface{} {
	objectMap := make(map[string]interface{})
	for s, v := range n.objectMap {
		objectMap[s.Evaluate().(string)] = v.Evaluate()
	}
	return objectMap
}
