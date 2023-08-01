package service

import (
	"context"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/app/task/job/biz"
)

type CometService struct {
	v1.UnimplementedCometAPIServer
	uc *biz.CometUseCase
}

func NewCometService(uc *biz.CometUseCase) *CometService {
	return &CometService{uc: uc}
}

func (t *CometService) Dispatch(ctx context.Context, in *v1.CometDispatchRequest) (out *v1.CometDispatchResponse, err error) {
	err = t.uc.Dispatch(ctx, in)
	if err != nil {
		return
	}
	out = &v1.CometDispatchResponse{}
	return
}
func (t *CometService) DispatchAsync(ctx context.Context, in *v1.CometDispatchRequest) (out *v1.CometDispatchResponse, err error) {
	err = t.uc.Dispatch(ctx, in)
	if err != nil {
		return
	}
	out = &v1.CometDispatchResponse{}
	return
}
