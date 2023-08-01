package data

import (
	"context"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/task/job/biz"
	"rockimserver/app/task/job/data/grpc/comet"
	"rockimserver/app/task/job/data/mq"
)

type cometRepo struct {
	grpc *comet.Manager
	mq   *mq.CometMessageManager
}

func NewCometRepo(grpc *comet.Manager, mq *mq.CometMessageManager) biz.CometRepo {
	return &cometRepo{grpc: grpc, mq: mq}
}

func (r *cometRepo) Dispatch(ctx context.Context, message *types.CometMessage) error {
	return r.grpc.Dispatch(ctx, message)
}

func (r *cometRepo) DispatchAsync(ctx context.Context, message *types.CometMessage) error {
	return r.mq.SaveMessage(ctx, message)
}
