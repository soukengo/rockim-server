package server

import (
	"context"
	"github.com/soukengo/gopkg/component/transport/queue"
	"github.com/soukengo/gopkg/component/transport/queue/iface"
	"github.com/soukengo/gopkg/component/transport/queue/options"
	"github.com/soukengo/gopkg/log"
	"github.com/soukengo/gopkg/util/codec"
)

func handleAndAck[Request any, Response any](fn func(context.Context, *Request) (*Response, error), decoder codec.Decoder) *queue.Handler {
	return queue.HandleWithOptions(func(ctx context.Context, req *Request) queue.Action {
		_, err := fn(ctx, req)
		if err != nil {
			log.WithContext(ctx).Error("handle failed", log.Pairs{"err": err})
			return iface.ReconsumeLater
		}
		return iface.CommitMessage
	}, options.Consumer().SetMode(options.Async).SetDecoder(decoder))
}
