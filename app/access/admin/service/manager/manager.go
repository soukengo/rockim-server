package manager

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAuthService, NewSessionService, NewPlatUserService, NewPlatRoleService, NewPlatResourceService)
