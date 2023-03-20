package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewProductService,
	NewUserService,
	NewAuthService,
	NewChatRoomService,
	NewChatRoomMemberService,
	NewMessageService,
)
