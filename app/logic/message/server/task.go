package server

import (
	"context"
	"github.com/soukengo/gopkg/component/transport/queue"
	"github.com/soukengo/gopkg/component/transport/queue/iface"
	"github.com/soukengo/gopkg/component/transport/queue/options"
	"github.com/soukengo/gopkg/log"
	"github.com/soukengo/gopkg/util/codec"
	"rockimserver/app/logic/message/biz/consts"
	"rockimserver/app/logic/message/task"
)

type TaskRegistry struct {
	messageTask *task.MessageTask
}

func NewTaskRegistry(messageTask *task.MessageTask) *TaskRegistry {
	return &TaskRegistry{messageTask: messageTask}
}

func (g *TaskRegistry) RegisterQueue(srv queue.Server) {
	srv.Subscribe(consts.QueueMessageDelivery, handleAndAck(g.messageTask.Delivery, codec.JSON))
}

func handleAndAck[Request any](fn func(context.Context, *Request) error, decoder codec.Decoder) *queue.Handler {
	return queue.HandleWithOptions(func(ctx context.Context, req *Request) queue.Action {
		err := fn(ctx, req)
		if err != nil {
			log.WithContext(ctx).Error("handle failed", log.Pairs{"err": err})
			return iface.ReconsumeLater
		}
		return iface.CommitMessage
	}, options.Consumer().SetMode(options.Async).SetDecoder(decoder))
}
