package manager

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAuthUseCase,
	NewSessionUseCase,
	NewPlatUserUseCase,
	NewPlatRoleUseCase,
	NewPlatResourceUseCase,
	NewTenantUseCase,
	NewTenantResourceUseCase,
)
