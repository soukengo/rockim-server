package database

import (
	"github.com/google/wire"
	"rockim/pkg/component/database/mongo"
)

var ProviderSet = wire.NewSet(mongo.NewClient, NewPlatUserData, NewPlatRoleData, NewPlatResourceData, NewTenantData, NewTenantResourceData)
