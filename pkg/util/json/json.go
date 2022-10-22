package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var json = jsoniter.ConfigDefault

func init() {
	extra.RegisterFuzzyDecoders()
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
