package lock

import (
	"rockimserver/pkg/component/database/redis"
)

type redisBuilder struct {
	cli *redis.Client
}

func NewRedisBuilder(cli *redis.Client) Builder {
	return &redisBuilder{cli: cli}
}

func (b *redisBuilder) Build(key Key, parts ...string) DistributedLock {
	return NewRedisDistributedLock(b.cli, key.GenKey(parts...))
}
