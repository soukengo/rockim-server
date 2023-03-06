package module

import (
	"github.com/google/wire"
	"rockimserver/app/access/comet/module/client"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(client.ProviderSet)
