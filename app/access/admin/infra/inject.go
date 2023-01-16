package infra

import (
	"github.com/google/wire"
	"rockim/app/access/admin/infra/database"
	"rockim/app/access/admin/infra/grpc"
)

var ProviderSet = wire.NewSet(grpc.ProviderSet, database.ProviderSet)
