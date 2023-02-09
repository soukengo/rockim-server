package redis

import (
	"context"
	rds "github.com/go-redis/redis/v8"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/errors"
)

type sortedSetCache[T any] struct {
	redisCache[T]
}

func NewSortedSetCache[T any](client *redis.Client, key string, opts ...cache.Option) cache.SortedSetCache[T] {
	return newSortedSetCache[T](client, key, cache.Default().Apply(opts...))
}

func newSortedSetCache[T any](cli *redis.Client, key string, opts *cache.Options) cache.SortedSetCache[T] {
	return &sortedSetCache[T]{redisCache: newRedisCache[T](cli, key, opts)}
}

func (c *sortedSetCache[T]) Add(ctx context.Context, member *cache.SortedMember[T]) (err error) {
	data, err := c.encode(member.Value)
	if err != nil {
		return
	}
	_, err = c.cli.ZAdd(ctx, c.key, &rds.Z{Score: member.Score, Member: data})
	return
}

func (c *sortedSetCache[T]) AddSlice(ctx context.Context, values []*cache.SortedMember[T]) (err error) {
	var records = make([]any, 0)
	for i, v := range values {
		var data []byte
		data, err = c.encode(v)
		if err != nil {
			return
		}
		records[i] = data
	}
	_, err = c.cli.SAdd(ctx, c.key, records...)
	return
}

func (c *sortedSetCache[T]) Paginate(ctx context.Context, cursor uint64, count int64) (ret []*T, retCursor uint64, err error) {
	values, retCursor, err := c.cli.SScan(ctx, c.key, cursor, "*", count)
	if err != nil {
		return
	}
	if retCursor == 0 {
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
	ret = make([]*T, len(values))
	for i, v := range values {
		var item *T
		item, err = c.decode([]byte(v))
		if err != nil {
			return
		}
		ret[i] = item
	}
	return
}

func (c *sortedSetCache[T]) DeleteMember(ctx context.Context, v *T) (err error) {
	data, err := c.encode(v)
	if err != nil {
		return
	}
	_, err = c.cli.ZRem(ctx, c.key, data)
	return
}

func (c *sortedSetCache[T]) ExistsMember(ctx context.Context, v *T) (isMember bool, err error) {
	data, err := c.encode(v)
	if err != nil {
		return
	}
	_, err = c.cli.ZRank(ctx, c.key, string(data))
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
			return
		}
	}
	return
}

func (c *sortedSetCache[T]) Count(ctx context.Context) (count int64, err error) {
	count, err = c.cli.ZCard(ctx, c.key)
	if err != nil {
		return
	}
	if count == 0 {
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
