package redis

import (
	"context"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/errors"
)

type valueCache[T any] struct {
	redisCache[T]
}

func NewValueCache[T any](client *redis.Client, key string, opts ...cache.Option) cache.ValueCache[T] {
	return newValueCache[T](client, key, cache.Default().Apply(opts...))
}

func newValueCache[T any](cli *redis.Client, key string, opts *cache.Options) cache.ValueCache[T] {
	return &valueCache[T]{redisCache: newRedisCache[T](cli, key, opts)}
}

func (c *valueCache[T]) Get(ctx context.Context) (ret *T, err error) {
	v, err := c.cli.Get(ctx, c.key)
	if err != nil {
		if errors.IsNotFound(err) {
			err = cache.ErrNoCache
		}
		return
	}
	return c.decode([]byte(v))
}

func (c *valueCache[T]) Set(ctx context.Context, v *T) (err error) {
	data, err := c.encode(v)
	if err != nil {
		return
	}
	_, err = c.cli.Set(ctx, c.key, data, c.opts.Expire())
	return
}
