package data

import (
	"context"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/app/task/job/biz"
	"rockimserver/app/task/job/data/grpc"
)

type pushRepo struct {
	grpc *grpc.PushManager
}

func NewPushRepo(grpc *grpc.PushManager) biz.PushRepo {
	return &pushRepo{grpc: grpc}
}

func (r *pushRepo) Push(ctx context.Context, message *types.Message) error {
	return r.grpc.Push(ctx, message)
}
