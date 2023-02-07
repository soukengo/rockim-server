package redis

import (
	"context"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
)

type valueCache[T any] struct {
	cli  *redis.Client
	key  string
	opts *cache.Options
}

func NewValueCache[T any](client *redis.Client, key string, opts ...cache.Option) cache.ValueCache[T] {
	return newValueCache[T](client, key, cache.Default().Apply(opts...))
}

func newValueCache[T any](cli *redis.Client, key string, opts *cache.Options) cache.ValueCache[T] {
	return &valueCache[T]{cli: cli, key: key, opts: opts}
}

func (c *valueCache[T]) Exists(ctx context.Context) (bool, error) {
	exists, err := c.cli.Exists(ctx, c.key)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (c *valueCache[T]) Delete(ctx context.Context) (err error) {
	_, err = c.cli.Del(ctx, c.key)
	return
}

func (c *valueCache[T]) Get(ctx context.Context) (ret *T, err error) {
	str, err := c.cli.Get(ctx, c.key)
	if err != nil {
		return
	}
	var value = new(T)
	err = c.opts.Codec().Decode([]byte(str), value)
	if err != nil {
		return
	}
	return value, nil
}

func (c *valueCache[T]) Set(ctx context.Context, value *T) (err error) {
	data, err := c.opts.Codec().Encode(value)
	if err != nil {
		return
	}
	_, err = c.cli.Set(ctx, c.key, data, c.opts.Expire())
	return
}
