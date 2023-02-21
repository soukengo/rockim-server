package module

import (
	"github.com/google/wire"
	"rockimserver/app/access/gateway/module/client"
	"rockimserver/app/access/gateway/module/openapi"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(client.ProviderSet, openapi.ProviderSet)
