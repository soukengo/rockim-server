package data

import (
	"context"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/logic/message/biz"
	"rockimserver/app/logic/message/data/mq"
)

type pushMessageRepo struct {
	mq *mq.PushMessageData
}

func NewPushMessageRepo(mq *mq.PushMessageData) biz.PushMessageRepo {
	return &pushMessageRepo{mq: mq}
}

func (r *pushMessageRepo) SavePushMessage(ctx context.Context, messages []*types.CometMessage) error {
	return r.mq.SavePushMessage(ctx, messages)
}
