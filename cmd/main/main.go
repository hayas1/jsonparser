package main

import (
	"fmt"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func main() {
	example := []string{
		`{`,
		`    "jsonparser" : "json parser implemented by go",`,
		`    "version": 0.1,`,
		`    "keyword": ["json", "parser", "go"]`,
		`}`,
	}
	// example := []string{`  true`}
	// fmt.Println(example)
	parser := jp.NewParser(example)
	node, parseErr := parser.ParseValue()
	fmt.Println(node.Evaluate(), parseErr)
}
