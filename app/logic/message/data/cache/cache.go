package cache

import (
	"rockimserver/pkg/component/cache"
	rediscache "rockimserver/pkg/component/cache/redis"
	"rockimserver/pkg/component/database/redis"
)

const (
	keyMessage         cache.Key = "message"
	keyMessageBox      cache.Key = "message.box"
	keyMessageDelivery cache.Key = "message.delivery"
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
