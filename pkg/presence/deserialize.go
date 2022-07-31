package presence

import (
	"bufio"
	"io"
	"os"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
	"github.com/hayas1/jsonparser/pkg/jsonparser/ast"
)

func Deserialize(jsonString []string) (ast.ValueNode, error) {
	parser := jp.NewParser(jsonString)
	return parser.ParseValue()
}

func DeserializeRead(reader io.Reader) (ast.ValueNode, error) {
	jsonString := make([]string, 0)
	r := bufio.NewScanner(reader)
	for r.Scan() {
		jsonString = append(jsonString, r.Text())
	}

	return Deserialize(jsonString)
}

func DeserializeFile(jsonPath string) (ast.ValueNode, error) {
	fp, err := os.Open(jsonPath)
	if err != nil {
		return ast.ValueNode{}, err
	}
	defer fp.Close()

	return DeserializeRead(fp)
}
