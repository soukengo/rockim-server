package data

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTokenRepo,
	NewChannelRepo,
	NewRoomRepo,
)
