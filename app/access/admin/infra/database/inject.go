package database

import (
	"github.com/google/wire"
	"rockimserver/app/access/admin/conf"
	"rockimserver/pkg/component/database/mongo"
)

var ProviderSet = wire.NewSet(NewMongoClient)

func NewMongoClient(cfg *conf.Database) *mongo.Client {
	return mongo.NewClient(cfg.Mongodb)
}
