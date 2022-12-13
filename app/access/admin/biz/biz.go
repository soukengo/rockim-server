package biz

import (
	"github.com/google/wire"
	"rockim/app/access/admin/biz/manager"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(manager.ProviderSet)
