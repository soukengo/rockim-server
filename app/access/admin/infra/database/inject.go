package database

import (
	"github.com/google/wire"
	"rockimserver/pkg/component/database"
	"rockimserver/pkg/component/database/mongo"
)

var ProviderSet = wire.NewSet(NewMongoClient)

func NewMongoClient(cfg *database.Config) *mongo.Client {
	return mongo.NewClient(cfg.Mongodb)
}
