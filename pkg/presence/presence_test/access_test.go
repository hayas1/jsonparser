package presence_test

import (
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/presence"
)

func TestValidAccess(tester *testing.T) {
	js := []string{
		`{`,
		`    "jsonparser" : "json parser implemented by go",`,
		`    "version": 0.1,`,
		`    "keyword": ["json", "parser", "go", {"one": 1, "two":2, "three" :3}]`,
		`}`,
	}
	root, err := jp.Deserialize(js)
	if err != nil {
		tester.Error("unexpected error: ", err)
	}

	if jp.Route(&root, jp.ObjInd("keyword"), jp.ArrInd(2)) != "go" {
		tester.Error("something go wrong")
	}
	if found, err := jp.Access(&root, "keyword", 3, "one"); !(found.(float64) == 1 && err == nil) {
		tester.Error("something go wrong")
	}
}

func TestInvalidAccess(tester *testing.T) {
	js := []string{
		`{`,
		`    "jsonparser" : "json parser implemented by go",`,
		`    "version": 0.1,`,
		`    "keyword": ["json", "parser", "go", {"one": 1, "two":2, "three" :3}]`,
		`}`,
	}
	root, err := jp.Deserialize(js)
	if err != nil {
		tester.Error("unexpected error: ", err)
	}

	if _, err := jp.Access(&root, "keyword", 3, "one", 1); err == nil {
		tester.Error("float64 cannot be indexed by integer 1")
	}
	if _, err := jp.Access(&root, "keyword", true); err != nil {
		tester.Error("weak type, so jp.Access will receive not int/string object...")
	}
}
