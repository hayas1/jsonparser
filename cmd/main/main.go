package main

import (
	"fmt"

	jp "github.com/hayas1/jsonparser/pkg/jsonparser"
)

func main() {
	emptyObject := []string{" \t \r\n[ \n \t \t ] \t\t"}
	parser := jp.NewParser(emptyObject)
	objectArray, err := parser.ParseArray()
	fmt.Println(objectArray, err)
	_, ok := objectArray.Evaluate().([]interface{})
	fmt.Println(ok)
}
