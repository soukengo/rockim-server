package redis

import (
	"context"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
)

type redisCache[T any] struct {
	cli  *redis.Client
	key  string
	opts *cache.Options
}

func newRedisCache[T any](cli *redis.Client, key string, opts *cache.Options) redisCache[T] {
	return redisCache[T]{cli: cli, key: key, opts: opts}
}

func (c *redisCache[T]) Exists(ctx context.Context) (bool, error) {
	exists, err := c.cli.Exists(ctx, c.key)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (c *redisCache[T]) Delete(ctx context.Context) (err error) {
	_, err = c.cli.Del(ctx, c.key)
	return
}

func (c *redisCache[T]) encode(v any) (data []byte, err error) {
	return c.opts.Codec().Encode(v)
}

func (c *redisCache[T]) decode(data []byte) (ret *T, err error) {
	var value = new(T)
	err = c.opts.Codec().Decode(data, value)
	if err != nil {
		return
	}
	return value, nil
}
