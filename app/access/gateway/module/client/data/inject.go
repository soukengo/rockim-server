package data

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewProductRepo,
	NewUserRepo,
	NewAuthRepo,
	NewChatRoomRepo,
	NewChatRoomMemberRepo,
	NewMessageRepo,
)
