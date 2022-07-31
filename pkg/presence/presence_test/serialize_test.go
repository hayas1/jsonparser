package presence_test

import (
	"reflect"
	"strings"
	"testing"

	jp "github.com/hayas1/jsonparser/pkg/presence"
)

func TestSerialize(tester *testing.T) {
	js := []string{
		`{`,
		`    "keyword": ["json", "parser", "go",`,
		`        "JavaScript Object Notation"`,
		`    ]`,
		`}`,
	}
	root, err := jp.Deserialize(js)
	if err != nil {
		tester.Error("unexpected error: ", err)
	}
	jStr := []string{
		`{`,
		`    "keyword": [`,
		`        "json",`,
		`        "parser",`,
		`        "go",`,
		`        "JavaScript Object Notation"`,
		`    ]`,
		`}`,
	}
	if jp.Serialize(&root) != strings.Join(jStr, "\n") {
		tester.Error(jp.Serialize(&root))
	}
}
func TestEquivalent(tester *testing.T) {
	js := []string{
		`{`,
		`    "jsonparser" : "json parser implemented by go",`,
		`    "version": 0.1,`,
		`    "keyword": ["json", "parser", "go",`,
		`        {"one": 1, "two":2, "three" :3}`,
		`    ]`,
		`}`,
	}
	root, err := jp.Deserialize(js)
	if err != nil {
		tester.Error("unexpected error: ", err)
	}

	js2 := jp.Serialize(&root)
	root2, err2 := jp.Deserialize([]string{js2})
	if err2 != nil {
		tester.Error("unexpected error: ", err2)
	}

	if !reflect.DeepEqual(root.Evaluate(), root2.Evaluate()) {
		tester.Error("equivalent error")
	}

}
