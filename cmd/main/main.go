package main

import (
	"fmt"

	"github.com/hayas1/jsonparser/pkg/jsonparser"
)

func main() {
	ch := make(chan jsonparser.Token)
	go jsonparser.Parse("json string", ch)
	fmt.Println(<-ch)
}
