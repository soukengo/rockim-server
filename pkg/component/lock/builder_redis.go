package lock

import (
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/log"
)

type redisBuilder struct {
	logger log.Logger
	cli    *redis.Client
}

func NewRedisBuilder(logger log.Logger, cli *redis.Client) Builder {
	return &redisBuilder{cli: cli, logger: logger}
}

func (b *redisBuilder) Build(key Key, parts ...string) DistributedLock {
	return NewRedisDistributedLock(b.logger, b.cli, key.GenKey(parts...))
}
