package database

import (
	"github.com/google/wire"
	"rockim/app/access/admin/conf"
	"rockim/pkg/component/database/mongo"
)

var ProviderSet = wire.NewSet(NewMongoClient)

func NewMongoClient(cfg *conf.Database) *mongo.Client {
	return mongo.NewClient(cfg.Mongodb)
}
