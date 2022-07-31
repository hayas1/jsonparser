package ast

import (
	"strconv"
	"strings"
)

type ObjectNode struct {
	ValueObject map[string]ValueNode
}

func (n *ObjectNode) Evaluate() interface{} {
	return n.Object()
}

func (n *ObjectNode) Dump(indent int) string {
	children := make([]string, 0)
	indentExternal := strings.Repeat(" ", 4*indent)
	indentInternal := strings.Repeat(" ", 4*(indent+1))
	for key, value := range n.ValueObject {
		children = append(children, indentInternal+strconv.Quote(key)+": "+value.Dump(indent+1))
	}
	return "{\n" + strings.Join(children, ",\n") + "\n" + indentExternal + "}"
}

func (n *ObjectNode) Object() map[string]interface{} {
	object := make(map[string]interface{})
	for s, v := range n.ValueObject {
		object[s] = v.Evaluate()
	}
	return object
}
