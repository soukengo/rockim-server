package mq

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/soukengo/gopkg/component/queue"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/logic/message/biz/consts"
)

type PushMessageData struct {
	cli queue.Producer
}

func NewPushMessageData(cli queue.Producer) *PushMessageData {
	return &PushMessageData{cli: cli}
}

func (d *PushMessageData) SavePushMessage(ctx context.Context, messages []*types.Message) (err error) {
	var req = &v1.MessagePushRequest{
		List: messages,
	}
	var bytes []byte
	bytes, err = proto.Marshal(req)
	if err != nil {
		return
	}
	err = d.cli.Submit(ctx, consts.DistributedQueueMessagePush, []queue.Value{bytes})
	if err != nil {
		return
	}
	return
}
