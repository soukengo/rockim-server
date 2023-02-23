package database

import (
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
)

type Config struct {
	Mongodb *mongo.Config
	Redis   *redis.Config
}
