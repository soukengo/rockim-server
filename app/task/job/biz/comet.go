package biz

import (
	"context"
	"github.com/soukengo/gopkg/log"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
)

type CometRepo interface {
	Dispatch(ctx context.Context, message *types.CometMessage) error
}

type CometUseCase struct {
	repo CometRepo
}

func NewMessagePushUseCase(repo CometRepo) *CometUseCase {
	return &CometUseCase{repo: repo}
}

func (uc *CometUseCase) Dispatch(ctx context.Context, in *v1.CometDispatchRequest) (err error) {
	for _, message := range in.List {
		err1 := uc.repo.Dispatch(ctx, message)
		if err1 != nil {
			log.WithContext(ctx).Errorf("Dispatch message error: %v", err1)
		}
	}
	return nil
}
