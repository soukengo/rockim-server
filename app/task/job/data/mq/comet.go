package mq

import (
	"context"
	"github.com/soukengo/gopkg/component/transport/queue"
	"github.com/soukengo/gopkg/component/transport/queue/options"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/task/job/biz/consts"
)

type CometMessageManager struct {
	srv queue.Server
}

func NewCometMessageManager(srv queue.Server) *CometMessageManager {
	return &CometMessageManager{srv: srv}
}

func (d *CometMessageManager) SaveMessage(ctx context.Context, message *types.CometMessage) (err error) {
	var req = &v1.CometDispatchRequest{
		Message: message,
	}

	err = d.srv.Publish(ctx, queue.NewRawMessage(consts.QueueCometDispatch, req), options.Producer())
	if err != nil {
		return
	}
	return
}
