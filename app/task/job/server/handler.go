package server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/soukengo/gopkg/component/server/job"
	"github.com/soukengo/gopkg/errors"
)

var (
	ErrInvalidData = errors.BadRequest("INVALID_DATA", "invalid data")
)

func wrapAction[Request any, Response any](fn func(context.Context, *Request) (*Response, error)) job.Handler {
	return func(ctx context.Context, topic string, data []byte) (err error) {
		var req = new(Request)
		err = decode(data, req)
		if err != nil {
			return
		}
		_, err = fn(ctx, req)
		return
	}
}

func decode(data []byte, v any) error {
	in, ok := v.(proto.Message)
	if !ok {
		return ErrInvalidData
	}
	return proto.Unmarshal(data, in)
}

func encode(v any) ([]byte, error) {
	in, ok := v.(proto.Message)
	if !ok {
		return nil, ErrInvalidData
	}
	return proto.Marshal(in)
}
