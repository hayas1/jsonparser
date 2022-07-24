package jsonparser_test

import (
	"testing"

	"github.com/hayas1/jsonparser/pkg/jsonparser"
)

func TestTokenizeBrace(tester *testing.T) {
	leftbrace, rightbrace := "{", "}"
	tl, tr := jsonparser.Tokenize(leftbrace), jsonparser.Tokenize(rightbrace)
	if !(tl.Element == "{" && tr.Element == "}") {
		tester.Error("brace's element should be \"{\" or \"}\"")
	}
}
