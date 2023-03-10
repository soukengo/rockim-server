package service

import "github.com/google/wire"

// ProviderSet is db providers.
var ProviderSet = wire.NewSet(
	NewMessageService,
)
