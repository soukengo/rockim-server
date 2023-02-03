package infra

import (
	"github.com/google/wire"
	"rockimserver/app/access/admin/infra/database"
	"rockimserver/app/access/admin/infra/grpc"
)

var ProviderSet = wire.NewSet(grpc.ProviderSet, database.ProviderSet)
