# jsonparser
jsonparser is simple json parser implemented by Go.
there may be a little few bugs remain.

## usage
### from/to file
```go
package main
import (
	"fmt"
	jp "github.com/hayas1/jsonparser/pkg/presence"
)

func main() {
	// `path/to.json`
	//
	// {
	//     "jsonparser": "json parser implemented by go",
	//     "version": 0.1,
	//     "keyword": [
	//         "json",
	//         "parser",
	//         "go",
	//         {
	//             "one": 1,
	//             "two": 2,
	//             "three": 3
	//         }
	//     ]
	// }

	// read from file
	root, _ := jp.DeserializeFile("path/to/read.json")

	// access to value
	fmt.Println(jp.Route(&root, jp.ObjInd("keyword"), jp.ArrInd(2))) // go
	fmt.Println(jp.Access(&root, "keyword", 3, "one"))               // 1 <nil>

	// write to file
	jp.SerializeFile("path/to/write.json", &root)
}
```

### from/to string
```go
package main
import (
	"fmt"
	jp "github.com/hayas1/jsonparser/pkg/presence"
)

func main() {
	js := []string{
		`{`,
		`    "jsonparser" : "json parser implemented by go",`,
		`    "version": 0.1,`,
		`    "keyword": ["json", "parser", "go",`,
		`        {"one": 1, "two":2, "three" :3}`,
		`    ]`,
		`}`,
	}

	// read json
	root, _ := jp.Deserialize(js)

	// access to value
	fmt.Println(jp.Route(&root, jp.ObjInd("keyword"), jp.ArrInd(2))) // go
	fmt.Println(jp.Access(&root, "keyword", 3, "one"))               // 1 <nil>

	// write json
	fmt.Print(jp.Serialize(&root))
	// {
	//     "jsonparser": "json parser implemented by go",
	//     "version": 0.1,
	//     "keyword": [
	//         "json",
	//         "parser",
	//         "go",
	//         {
	//             "one": 1,
	//             "two": 2,
	//             "three": 3
	//         }
	//     ]
	// }
}
```