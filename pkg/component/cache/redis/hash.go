package redis

import (
	"context"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/errors"
)

type hashCache[T any] struct {
	redisCache[T]
}

func NewHashCache[T any](client *redis.Client, key string, opts ...cache.Option) cache.HashCache[T] {
	return newHashCache[T](client, key, cache.Default().Apply(opts...))
}

func newHashCache[T any](cli *redis.Client, key string, opts *cache.Options) cache.HashCache[T] {
	return &hashCache[T]{redisCache: newRedisCache[T](cli, key, opts)}
}

func (c *hashCache[T]) Put(ctx context.Context, field string, v *T) (err error) {
	data, err := c.encode(v)
	if err != nil {
		return
	}
	_, err = c.cli.HSet(ctx, c.key, field, data)
	return
}

func (c *hashCache[T]) PutMap(ctx context.Context, hash map[string]*T) (err error) {
	var records = make(map[string]any)
	for k, v := range hash {
		var data []byte
		data, err = c.encode(v)
		if err != nil {
			return
		}
		records[k] = data
	}
	_, err = c.cli.HMSet(ctx, c.key, records)
	return
}

func (c *hashCache[T]) Get(ctx context.Context, field string) (ret *T, err error) {
	v, err := c.cli.HGet(ctx, c.key, field)
	if err != nil {
		if errors.IsNotFound(err) {
			var exists bool
			exists, err = c.Exists(ctx)
			if err != nil {
				return
			}
			if !exists {
				err = cache.ErrNoCache
				return
			}
		}
		return
	}
	return c.decode([]byte(v))
}

func (c *hashCache[T]) GetAll(ctx context.Context) (ret map[string]*T, err error) {
	hash, err := c.cli.HGetAll(ctx, c.key)
	if err != nil {
		if errors.IsNotFound(err) {
			err = cache.ErrNoCache
		}
		return
	}
	ret = make(map[string]*T)
	for k, v := range hash {
		var item *T
		item, err = c.decode([]byte(v))
		if err != nil {
			return
		}
		ret[k] = item
	}
	return
}

func (c *hashCache[T]) ExistsField(ctx context.Context, field string) (existsField bool, err error) {
	existsField, err = c.cli.HExists(ctx, c.key, field)
	if err != nil {
		return
	}
	if !existsField {
		var exists bool
		exists, err = c.Exists(ctx)
		if err != nil {
			return
		}
		if !exists {
			err = cache.ErrNoCache
			return
		}
	}
	return
}

func (c *hashCache[T]) DeleteField(ctx context.Context, field string) (err error) {
	_, err = c.cli.HDel(ctx, c.key, field)
	return
}
