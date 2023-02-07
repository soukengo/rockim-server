package cache

import "rockimserver/pkg/util/encoding/json"

type Codec interface {
	Encode(any) ([]byte, error)
	Decode([]byte, any) error
}

type Primitive interface {
	~int32 | ~int64 | ~int | ~string
}

type jsonCodec struct {
}

func NewJsonCodec() Codec {
	return &jsonCodec{}
}

func (c *jsonCodec) Encode(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (c *jsonCodec) Decode(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
