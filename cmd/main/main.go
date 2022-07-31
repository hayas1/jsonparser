package main

import (
	"fmt"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func main() {
	example := []string{`"json\r\n parser\t\"g\u00f3\"😶"`, ``}
	fmt.Println(example)
	parser := jp.NewParser(example)
	node, parseErr := parser.ParseString()
	fmt.Println(node.Evaluate(), parseErr)
}
