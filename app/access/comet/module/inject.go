package module

import (
	"github.com/google/wire"
	"rockimserver/app/access/comet/module/client"
	"rockimserver/app/access/comet/module/server"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(client.ProviderSet, server.ProviderSet)
