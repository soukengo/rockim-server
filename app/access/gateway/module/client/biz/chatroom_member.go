package biz

import (
	"context"
	"rockimserver/app/access/gateway/module/client/biz/options"
)

type ChatRoomMemberRepo interface {
	Join(ctx context.Context, opts *options.ChatRoomJoinOptions) error
	Quit(ctx context.Context, opts *options.ChatRoomQuitOptions) error
}

type ChatRoomMemberUseCase struct {
	repo ChatRoomMemberRepo
}

func NewChatRoomMemberUseCase(repo ChatRoomMemberRepo) *ChatRoomMemberUseCase {
	return &ChatRoomMemberUseCase{repo: repo}
}

func (uc *ChatRoomMemberUseCase) Join(ctx context.Context, opts *options.ChatRoomJoinOptions) (err error) {
	return uc.repo.Join(ctx, opts)
}
func (uc *ChatRoomMemberUseCase) Quit(ctx context.Context, opts *options.ChatRoomQuitOptions) (err error) {
	return uc.repo.Quit(ctx, opts)
}
