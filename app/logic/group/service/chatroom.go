package service

import "rockimserver/app/logic/group/biz"

type ChatRoomService struct {
	uc *biz.ChatRoomUseCase
}

func NewChatRoomService(uc *biz.ChatRoomUseCase) *ChatRoomService {
	return &ChatRoomService{uc: uc}
}
