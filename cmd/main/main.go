package main

import (
	"fmt"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func main() {
	// example := []string{`{"json parser" : "implemented by go", "version": 0.1}`}
	example := []string{`true`}
	// fmt.Println(example)
	parser := jp.NewParser(example)
	node, parseErr := parser.ParseImmediate()
	fmt.Println(node.Evaluate(), parseErr)
}
