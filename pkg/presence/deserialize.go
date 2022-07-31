package presence

import (
	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
	"github.com/hayas1/jsonparser/pkg/jsonparser/ast"
)

func Deserialize(jsonString []string) (ast.ValueNode, error) {
	parser := jp.NewParser(jsonString)
	return parser.ParseValue()
}
