package presence_test

import (
	"reflect"
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/presence"
)

func TestFileRead(tester *testing.T) {
	root, parseErr := jp.DeserializeFile("test/simple.json")
	if parseErr != nil {
		tester.Error("unexpected error: ", parseErr)
	}

	v, err := jp.Access(&root, "version")
	if val, ok := v.(float64); !(err == nil && ok && val == 0.1) {
		tester.Error("something go wrong")
	}
}

func TestFileWrite(tester *testing.T) {
	root1, parseErr1 := jp.DeserializeFile("test/unique1.json")
	if parseErr1 != nil {
		tester.Error("unexpected error: ", parseErr1)
	}

	jp.SerializeFile("test/unique2.json", &root1)

	root2, parseErr2 := jp.DeserializeFile("test/unique2.json")
	if parseErr2 != nil {
		tester.Error("unexpected error: ", parseErr2)
	}

	if !reflect.DeepEqual(root1.Evaluate(), root2.Evaluate()) {
		tester.Error("something go wrong")
	}
}
