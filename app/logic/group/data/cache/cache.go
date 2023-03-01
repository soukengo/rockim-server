package cache

import (
	"rockimserver/pkg/component/cache"
	rediscache "rockimserver/pkg/component/cache/redis"
	"rockimserver/pkg/component/database/redis"
)

const (
	keyGroup              cache.Key = "group"
	keyGroupCustomGroupId cache.Key = "group.custom_group_id"

	keyGroupMemberSet  cache.Key = "group.member.set"
	keyGroupMemberInfo cache.Key = "group.member.info"

	keyChatRoomMemberSet  cache.Key = "chatroom.member.set"
	keyChatRoomMemberHash cache.Key = "chatroom.member.hash"
)

// todo golang 不支持 method type parameters，先这样写

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
