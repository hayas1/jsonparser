package presence

import (
	"github.com/hayas1/jsonparser/pkg/jsonparser/ast"
)

func Serialize(jsonRoot ast.AstNode) string {
	string := jsonRoot.Dump(0)
	return string
}
