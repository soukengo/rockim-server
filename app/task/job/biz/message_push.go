package biz

import (
	"context"
	v1 "rockimserver/apis/rockim/service/job/v1"
	"rockimserver/apis/rockim/service/job/v1/types"
	"rockimserver/pkg/log"
)

type PushRepo interface {
	Push(ctx context.Context, message *types.Message) error
}

type MessagePushUseCase struct {
	repo PushRepo
}

func NewMessagePushUseCase(repo PushRepo) *MessagePushUseCase {
	return &MessagePushUseCase{repo: repo}
}

func (uc *MessagePushUseCase) Push(ctx context.Context, in *v1.MessagePushRequest) (err error) {
	for _, message := range in.List {
		err1 := uc.repo.Push(ctx, message)
		if err1 != nil {
			log.WithContext(ctx).Errorf("Push message error: %v", err1)
		}
	}
	return nil
}
