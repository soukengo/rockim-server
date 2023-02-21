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

func NewSortedSetCache[T any](client *redis.Client, category *cache.Category, opts ...cache.Option) cache.SortedSetCache[T] {
	return newSortedSetCache[T](client, category, cache.Apply(category, opts...))
}

func newSortedSetCache[T any](cli *redis.Client, category *cache.Category, opts *cache.Options) cache.SortedSetCache[T] {
	return &sortedSetCache[T]{redisCache: newRedisCache[T](cli, category, opts)}
}

func (c *sortedSetCache[T]) Add(ctx context.Context, keySuffix string, member *cache.SortedMember[T]) (err error) {
	data, err := c.encode(member.Value)
	if err != nil {
		return
	}
	_, err = c.cli.ZAdd(ctx, c.key(keySuffix), &rds.Z{Score: member.Score, Member: data})
	return
}

func (c *sortedSetCache[T]) AddSlice(ctx context.Context, keySuffix string, values []*cache.SortedMember[T]) (err error) {
	var records = make([]any, 0)
	for i, v := range values {
		var data []byte
		data, err = c.encode(v)
		if err != nil {
			return
		}
		records[i] = data
	}
	_, err = c.cli.SAdd(ctx, c.key(keySuffix), records...)
	return
}

func (c *sortedSetCache[T]) Paginate(ctx context.Context, keySuffix string, cursor uint64, count int64) (ret []*T, retCursor uint64, err error) {
	values, retCursor, err := c.cli.SScan(ctx, c.key(keySuffix), cursor, "*", count)
	if err != nil {
		return
	}
	if retCursor == 0 {
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
	ret = make([]*T, len(values))
	for i, v := range values {
		var item *T
		item, err = c.decodeStr(v)
		if err != nil {
			return
		}
		ret[i] = item
	}
	return
}

func (c *sortedSetCache[T]) DeleteMember(ctx context.Context, keySuffix string, v *T) (err error) {
	data, err := c.encode(v)
	if err != nil {
		return
	}
	_, err = c.cli.ZRem(ctx, c.key(keySuffix), data)
	return
}

func (c *sortedSetCache[T]) ExistsMember(ctx context.Context, keySuffix string, v *T) (isMember bool, err error) {
	data, err := c.encode(v)
	if err != nil {
		return
	}
	_, err = c.cli.ZRank(ctx, c.key(keySuffix), string(data))
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
			return
		}
	}
	return
}

func (c *sortedSetCache[T]) Count(ctx context.Context, keySuffix string) (count int64, err error) {
	count, err = c.cli.ZCard(ctx, c.key(keySuffix))
	if err != nil {
		return
	}
	if count == 0 {
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
