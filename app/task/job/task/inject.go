package task

import "github.com/google/wire"

// ProviderSet is task providers.
var ProviderSet = wire.NewSet(
	NewMessageTask,
)
