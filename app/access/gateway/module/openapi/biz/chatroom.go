package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/access/gateway/module/openapi/biz/options"
)

type ChatRoomRepo interface {
	Create(ctx context.Context, opts *options.ChatRoomCreateOptions) (out *types.Group, err error)
}

type ChatRoomUseCase struct {
	repo ChatRoomRepo
}

func NewChatRoomUseCase(repo ChatRoomRepo) *ChatRoomUseCase {
	return &ChatRoomUseCase{repo: repo}
}

func (uc *ChatRoomUseCase) Create(ctx context.Context, opts *options.ChatRoomCreateOptions) (out *types.Group, err error) {
	return uc.repo.Create(ctx, opts)
}
