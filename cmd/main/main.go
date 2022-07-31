package main

import (
	"fmt"

	jp "github.com/hayas1/jsonparser/pkg/presence"
)

func main() {
	example := []string{
		`{`,
		`    "jsonparser" : "json parser implemented by go",`,
		`    "version": 0.1,`,
		`    "keyword": ["json", "parser", "go", {"one": 1, "two":2, "three" :3}]`,
		`}`,
	}
	root, err := jp.Deserialize(example)
	fmt.Println(root.Evaluate(), err) // parse result
	fmt.Println(jp.Serialize(&root))  // dump json

	fmt.Println(jp.Access(&root, jp.ObjInd("keyword"), jp.ArrInd(2))) // go
	fmt.Println(jp.Route(&root, "keyword", 3, "one"))                 // 1
}
