package data

import (
	"github.com/google/wire"
	"rockimserver/app/access/comet/module/server/data/socket"
)

var ProviderSet = wire.NewSet(socket.ProviderSet, NewChannelRepo)
