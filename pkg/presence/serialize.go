package presence

import (
	"bufio"
	"io"
	"os"

	"github.com/hayas1/jsonparser/pkg/jsonparser/ast"
)

func Serialize(jsonRoot ast.AstNode) string {
	string := jsonRoot.Dump(0)
	return string
}

func SerializeWrite(writer io.Writer, jsonRoot ast.AstNode) error {
	w := bufio.NewWriter(writer)
	if _, err := w.WriteString(Serialize(jsonRoot)); err != nil {
		return err
	}
	w.Flush()
	return nil
}

func SerializeFile(jsonPath string, jsonRoot ast.AstNode) error {
	fp, err := os.Create(jsonPath)
	if err != nil {
		return err
	}
	defer fp.Close()

	return SerializeWrite(fp, jsonRoot)
}
