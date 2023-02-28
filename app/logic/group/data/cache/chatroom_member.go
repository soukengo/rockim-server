package cache

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
)

type ChatRoomMemberData struct {
	cache    cache.ValueCache[types.GroupMember]
	setCache cache.SortedSetCache[string]
}

func NewChatRoomMemberData(rds *redis.Client, cfg *cache.Config) *ChatRoomMemberData {
	return &ChatRoomMemberData{
		cache:    newValueCache[types.GroupMember](rds, cfg.Category(keyChatRoomMemberInfo)),
		setCache: newSortedSetCache[string](rds, cfg.Category(keyChatRoomMemberSet)),
	}
}

func (d *ChatRoomMemberData) Save(ctx context.Context, member *types.GroupMember) (err error) {
	err = d.cache.Set(ctx, cache.Parts(member.ProductId, member.GroupId, member.Uid), member)
	if err != nil {
		return
	}
	setItem := &member.Uid
	return d.setCache.Add(ctx, cache.Parts(member.ProductId, member.GroupId), cache.Member[string](member.CreateTime, setItem))
}

func (d *ChatRoomMemberData) Delete(ctx context.Context, productId string, groupId string, uid string) (err error) {
	err = d.cache.Delete(ctx, cache.Parts(productId, groupId, uid))
	if err != nil {
		return
	}
	setItem := &uid
	return d.setCache.DeleteMember(ctx, cache.Parts(productId, groupId), setItem)
}

func (d *ChatRoomMemberData) Exists(ctx context.Context, productId string, groupId string, uid string) (exists bool, err error) {
	return d.cache.Exists(ctx, cache.Parts(productId, groupId, uid))
}
