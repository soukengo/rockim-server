package cache

import (
	"rockimserver/pkg/component/cache"
	rediscache "rockimserver/pkg/component/cache/redis"
	"rockimserver/pkg/component/database/redis"
	"strings"
)

const keySplit = ":"

const (
	keyUser             = "user"
	keyUserAccount      = "user.account"
	keyAuthCode         = "auth.code"
	keyAuthCodeReverse  = "auth.code.reverse"
	keyAuthToken        = "auth.token"
	keyAuthTokenReverse = "auth.token.reverse"
)

func genKey(keys ...string) string {
	return strings.Join(keys, keySplit)
}

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
