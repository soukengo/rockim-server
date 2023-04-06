package database

import (
	"github.com/google/wire"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/database/mongo"
)

var ProviderSet = wire.NewSet(NewMongoClient)

func NewMongoClient(cfg *database.Config) *mongo.Client {
	return mongo.NewClient(cfg.Mongodb)
}
