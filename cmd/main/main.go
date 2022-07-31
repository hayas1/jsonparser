package main

import (
	"fmt"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func main() {
	example := []string{`3e-1`, `.1`}
	// fmt.Println(example)
	parser := jp.NewParser(example)
	node, parseErr := parser.ParseNumber()
	fmt.Println(node.Evaluate(), parseErr)
}
