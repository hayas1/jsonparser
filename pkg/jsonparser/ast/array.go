package ast

import "strings"

type ArrayNode struct {
	ValueArray []ValueNode
}

func (n *ArrayNode) Evaluate() interface{} {
	return n.Array()
}

func (t *ArrayNode) Dump(indent int) string {
	children := make([]string, len(t.ValueArray))
	indentExternal := strings.Repeat(" ", 4*indent)
	indentInternal := strings.Repeat(" ", 4*(indent+1))
	for i := 0; i < len(t.ValueArray); i++ {
		children[i] = indentInternal + t.ValueArray[i].Dump(indent+1)
	}
	return "[\n" + strings.Join(children, ",\n") + "\n" + indentExternal + "]"
}

func (n *ArrayNode) Array() []interface{} {
	array := make([]interface{}, len(n.ValueArray))
	for i := 0; i < len(n.ValueArray); i++ {
		array[i] = n.ValueArray[i].Evaluate()
	}
	return array
}
