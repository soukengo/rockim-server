package biz

import (
	"context"
	"rockimserver/app/access/comet/module/server/biz/options"
)

type ChannelRepo interface {
	Push(ctx context.Context, opts *options.PushOptions) error
	PushRoom(ctx context.Context, opts *options.PushRoomOptions) error
}

type ChannelUseCase struct {
	repo ChannelRepo
}

func NewChannelUseCase(repo ChannelRepo) *ChannelUseCase {
	return &ChannelUseCase{repo: repo}
}

func (uc *ChannelUseCase) Push(ctx context.Context, opts *options.PushOptions) error {
	return uc.repo.Push(ctx, opts)
}

func (uc *ChannelUseCase) PushRoom(ctx context.Context, opts *options.PushRoomOptions) error {
	return uc.repo.PushRoom(ctx, opts)
}
