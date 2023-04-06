package cache

import (
	"github.com/soukengo/gopkg/component/cache"
	rediscache "github.com/soukengo/gopkg/component/cache/redis"
	"github.com/soukengo/gopkg/component/database/redis"
)

const (
	keyProduct cache.Key = "product"
)

func newValueCache[T any](client *redis.Client, category *cache.Category, opts ...cache.Option) cache.ValueCache[T] {
	return rediscache.NewValueCache[T](client, category, opts...)
}
func newHashCache[T any](client *redis.Client, category *cache.Category, opts ...cache.Option) cache.HashCache[T] {
	return rediscache.NewHashCache[T](client, category, opts...)
}
func newSetCache[T any](client *redis.Client, category *cache.Category, opts ...cache.Option) cache.SetCache[T] {
	return rediscache.NewSetCache[T](client, category, opts...)
}
func newSortedSetCache[T any](client *redis.Client, category *cache.Category, opts ...cache.Option) cache.SortedSetCache[T] {
	return rediscache.NewSortedSetCache[T](client, category, opts...)
}
