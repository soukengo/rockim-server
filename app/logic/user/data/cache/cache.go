package cache

import (
	"rockimserver/pkg/component/cache"
	rediscache "rockimserver/pkg/component/cache/redis"
	"rockimserver/pkg/component/database/redis"
)

func newValueCache[T any](client *redis.Client, key string, opts ...cache.Option) cache.ValueCache[T] {
	return rediscache.NewValueCache[T](client, key, opts...)
}
func newHashCache[T any](client *redis.Client, key string, opts ...cache.Option) cache.HashCache[T] {
	return rediscache.NewHashCache[T](client, key, opts...)
}
func newSetCache[T any](client *redis.Client, key string, opts ...cache.Option) cache.SetCache[T] {
	return rediscache.NewSetCache[T](client, key, opts...)
}
func newSortedSetCache[T any](client *redis.Client, key string, opts ...cache.Option) cache.SortedSetCache[T] {
	return rediscache.NewSortedSetCache[T](client, key, opts...)
}
