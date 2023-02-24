package biz

import "context"

type ChatRoomRepo interface {
	AddMember(ctx context.Context) error
}

type ChatRoomUseCase struct {
	repo ChatRoomRepo
}

func NewChatRoomUseCase(repo ChatRoomRepo) *ChatRoomUseCase {
	return &ChatRoomUseCase{repo: repo}
}
