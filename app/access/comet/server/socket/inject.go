package socket

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSocketServer)
