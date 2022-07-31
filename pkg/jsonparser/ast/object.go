package ast

type ObjectNode struct {
	ValueObject map[StringNode]ValueNode
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
