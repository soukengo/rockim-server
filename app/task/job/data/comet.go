package data

import (
	"context"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/task/job/biz"
	"rockimserver/app/task/job/data/grpc/comet"
)

type cometRepo struct {
	comet *comet.Manager
}

func NewCometRepo(grpc *comet.Manager) biz.CometRepo {
	return &cometRepo{comet: grpc}
}

func (r *cometRepo) Dispatch(ctx context.Context, message *types.CometMessage) error {
	return r.comet.Dispatch(ctx, message)
}
