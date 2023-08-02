package biz

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/idgen"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	idgen.NewMongoGenerator,
	NewTenantUseCase,
	NewProductUseCase,
)
