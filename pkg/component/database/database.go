package database

import (
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/component/database/redis"
)

func NewRedisClient(cfg *Config) *redis.Client {
	return redis.NewClient(cfg.Redis)
}
func NewMongoClient(cfg *Config) *mongo.Client {
	return mongo.NewClient(cfg.Mongodb)
}
