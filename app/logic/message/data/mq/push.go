package mq

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/soukengo/gopkg/component/mq"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

type PushMessageData struct {
	cli mq.Producer
}

func NewPushMessageData(cli mq.Producer) *PushMessageData {
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
	err = d.cli.Send(ctx, enums.MQ_MESSAGE_PUSH.String(), bytes)
	if err != nil {
		return
	}
	return
}
