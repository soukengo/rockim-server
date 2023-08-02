package database

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database"
	"rockimserver/app/logic/platform/conf"
)

var ProviderSet = wire.NewSet(
	NewDatabaseManager,
	NewTenantData,
	NewProductData,
)

func NewDatabaseManager(cfg *conf.Config) *database.Manager {
	return database.NewManager(cfg.Database)
}
