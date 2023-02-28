package biz

import "context"

type ChatRoomMemberRepo interface {
	AddMember(ctx context.Context) error
}

type ChatRoomMemberUseCase struct {
	repo ChatRoomMemberRepo
}

func NewChatRoomMemberUseCase(repo ChatRoomMemberRepo) *ChatRoomMemberUseCase {
	return &ChatRoomMemberUseCase{repo: repo}
}
