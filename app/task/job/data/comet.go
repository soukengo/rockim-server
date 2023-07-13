package data

import (
	"context"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/task/job/biz"
	"rockimserver/app/task/job/data/grpc"
)

type cometRepo struct {
	grpc *grpc.CometManager
}

func NewCometRepo(grpc *grpc.CometManager) biz.CometRepo {
	return &cometRepo{grpc: grpc}
}

func (r *cometRepo) Dispatch(ctx context.Context, message *types.CometMessage) error {
	return r.grpc.Dispatch(ctx, message)
}
