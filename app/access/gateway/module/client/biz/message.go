package biz

import (
	"context"
	"rockimserver/app/access/gateway/module/client/biz/options"
)

type MessageRepo interface {
	Send(ctx context.Context, opts *options.MessageSendOptions) error
}

type MessageUseCase struct {
	repo MessageRepo
}

func NewMessageUseCase(repo MessageRepo) *MessageUseCase {
	return &MessageUseCase{repo: repo}
}

func (uc *MessageUseCase) Send(ctx context.Context, opts *options.MessageSendOptions) error {
	return uc.repo.Send(ctx, opts)
}
