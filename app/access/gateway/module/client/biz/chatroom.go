package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
)

type ChatRoomRepo interface {
	FindById(ctx context.Context, productId string, groupId string) (*types.Group, error)
	FindGroupId(ctx context.Context, productId string, customGroupId string) (string, error)
}

type ChatRoomUseCase struct {
	repo ChatRoomRepo
}

func NewChatRoomUseCase(repo ChatRoomRepo) *ChatRoomUseCase {
	return &ChatRoomUseCase{repo: repo}
}

func (uc *ChatRoomUseCase) Find(ctx context.Context, productId string, groupId string) (*types.Group, error) {
	return uc.repo.FindById(ctx, productId, groupId)
}
func (uc *ChatRoomUseCase) FindGroupId(ctx context.Context, productId string, customGroupId string) (string, error) {
	return uc.repo.FindGroupId(ctx, productId, customGroupId)
}
