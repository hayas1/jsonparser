package ast

type StringNode struct {
	Str string
}

func (n *StringNode) Evaluate() interface{} {
	return n.String()
}

func (n *StringNode) String() string {
	return n.Str
}
