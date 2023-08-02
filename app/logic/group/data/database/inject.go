package database

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database"
	"rockimserver/app/logic/group/conf"
)

var ProviderSet = wire.NewSet(
	NewDatabaseManager,
	NewChatRoomData,
)

func NewDatabaseManager(cfg *conf.Config) *database.Manager {
	return database.NewManager(cfg.Database)
}
