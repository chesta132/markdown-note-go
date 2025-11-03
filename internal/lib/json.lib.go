package lib

import (
	"encoding/json"
)

type jsonPkg struct{}

// Json is the global instance of jsonPkg for convenience.
var Json = &jsonPkg{}

// Stringify converts any Go data structure into a pretty-printed JSON string.
// Usage:
//
//	jsonStr := lib.Json.Stringify(myStruct)
func (*jsonPkg) Stringify(data any) string {
	d, _ := json.MarshalIndent(data, "", "  ")
	return string(d)
}
