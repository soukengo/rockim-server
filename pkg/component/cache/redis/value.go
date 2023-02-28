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

func NewValueCache[T any](client *redis.Client, category *cache.Category, opts ...cache.Option) cache.ValueCache[T] {
	return newValueCache[T](client, category, cache.Apply(category, opts...))
}

func newValueCache[T any](cli *redis.Client, category *cache.Category, opts *cache.Options) cache.ValueCache[T] {
	return &valueCache[T]{redisCache: newRedisCache[T](cli, category, opts)}
}

func (c *valueCache[T]) Get(ctx context.Context, parts cache.KeyParts) (ret *T, err error) {
	v, err := c.cli.Get(ctx, c.key(parts))
	if err != nil {
		if errors.IsNotFound(err) {
			err = cache.ErrNoCache
		}
		return
	}
	ret, err = c.decodeStr(v)
	return
}

func (c *valueCache[T]) Set(ctx context.Context, parts cache.KeyParts, v *T) (err error) {
	data, err := c.encode(v)
	if err != nil {
		return
	}
	expire := c.opts.Expire()
	if len(data) == 0 {
		expire = c.opts.EmptyExpire()
	}
	_, err = c.cli.Set(ctx, c.key(parts), data, expire)
	return
}
