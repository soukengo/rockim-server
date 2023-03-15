package mq

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"rockimserver/apis/rockim/service/delivery/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/pkg/component/mq/kafka"
)

type PushMessageData struct {
	cli *kafka.Producer
}

func NewPushMessageData(cli *kafka.Producer) *PushMessageData {
	return &PushMessageData{cli: cli}
}

func (d *PushMessageData) SavePushMessage(ctx context.Context, messages []*types.PushMessage) (err error) {
	key := uuid.New().String()
	for _, message := range messages {
		var bytes []byte
		bytes, err = proto.Marshal(message)
		if err != nil {
			return
		}
		_, _, err = d.cli.SendMessage(enums.MQ_MESSAGE_DELIVERY.String(), key, bytes)
		if err != nil {
			return
		}
	}
	return
}
