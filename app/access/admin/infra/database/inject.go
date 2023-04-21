package database

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database"
	"rockimserver/app/access/admin/conf"
)

var ProviderSet = wire.NewSet(NewDatabaseManager)

func NewDatabaseManager(cfg *conf.Config) *database.Manager {
	return database.NewManager(cfg.Database)
}
