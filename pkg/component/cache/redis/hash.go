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

func NewHashCache[T any](client *redis.Client, category *cache.Category, opts ...cache.Option) cache.HashCache[T] {
	return newHashCache[T](client, category, cache.Apply(category, opts...))
}

func newHashCache[T any](cli *redis.Client, category *cache.Category, opts *cache.Options) cache.HashCache[T] {
	return &hashCache[T]{redisCache: newRedisCache[T](cli, category, opts)}
}

func (c *hashCache[T]) Put(ctx context.Context, keySuffix string, field string, v *T) (err error) {
	data, err := c.encode(v)
	if err != nil {
		return
	}
	_, err = c.cli.HSet(ctx, c.key(keySuffix), field, data)
	return
}

func (c *hashCache[T]) PutMap(ctx context.Context, keySuffix string, hash map[string]*T) (err error) {
	var records = make(map[string]any)
	for k, v := range hash {
		var data []byte
		data, err = c.encode(v)
		if err != nil {
			return
		}
		records[k] = data
	}
	_, err = c.cli.HMSet(ctx, c.key(keySuffix), records)
	return
}

func (c *hashCache[T]) Get(ctx context.Context, keySuffix string, field string) (ret *T, err error) {
	v, err := c.cli.HGet(ctx, c.key(keySuffix), field)
	if err != nil {
		if errors.IsNotFound(err) {
			var exists bool
			exists, err = c.Exists(ctx, keySuffix)
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
	return c.decodeStr(v)
}

func (c *hashCache[T]) GetAll(ctx context.Context, keySuffix string) (ret map[string]*T, err error) {
	hash, err := c.cli.HGetAll(ctx, c.key(keySuffix))
	if err != nil {
		if errors.IsNotFound(err) {
			err = cache.ErrNoCache
		}
		return
	}
	ret = make(map[string]*T)
	for k, v := range hash {
		var item *T
		item, err = c.decodeStr(v)
		if err != nil {
			return
		}
		ret[k] = item
	}
	return
}

func (c *hashCache[T]) ExistsField(ctx context.Context, keySuffix string, field string) (existsField bool, err error) {
	existsField, err = c.cli.HExists(ctx, c.key(keySuffix), field)
	if err != nil {
		return
	}
	if !existsField {
		var exists bool
		exists, err = c.Exists(ctx, keySuffix)
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

func (c *hashCache[T]) DeleteField(ctx context.Context, keySuffix string, field string) (err error) {
	_, err = c.cli.HDel(ctx, c.key(keySuffix), field)
	return
}
