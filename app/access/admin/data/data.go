package data

import (
	"github.com/google/wire"
	"rockim/app/access/admin/data/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(grpc.ProviderSet, NewPlatUserRepo, NewPlatRoleRepo, NewPlatResourceRepo, NewManagerTenantRepo, NewManagerTenantResourceRepo, NewTenantRepo, NewTenantResourceRepo)
