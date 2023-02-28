package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
)

type ChatRoomMemberRepo interface {
	Add(ctx context.Context, member *types.GroupMember) error
	Update(ctx context.Context, member *types.GroupMember) error
	Delete(ctx context.Context, productId string, groupId string, uid string) error
	Exists(ctx context.Context, productId string, groupId string, uid string) (bool, error)
}

type ChatRoomMemberUseCase struct {
	repo ChatRoomMemberRepo
}

func NewChatRoomMemberUseCase(repo ChatRoomMemberRepo) *ChatRoomMemberUseCase {
	return &ChatRoomMemberUseCase{repo: repo}
}
