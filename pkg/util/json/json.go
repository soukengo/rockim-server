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

func ToJSONString(v interface{}) (result string, err error) {
	if v == nil {
		return
	}
	result, err = json.MarshalToString(v)
	return
}
func TryToJSONString(v interface{}) (result string) {
	result, _ = ToJSONString(v)
	return
}
