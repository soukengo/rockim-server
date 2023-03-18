package database

import (
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/log"
)

func NewRedisClient(cfg *Config, logger log.Logger) *redis.Client {
	return redis.NewClient(cfg.Redis, logger)
}
func NewMongoClient(cfg *Config, logger log.Logger) *mongo.Client {
	return mongo.NewClientWithOptions(cfg.Mongodb, mongo.WithLogger(logger))
}
