package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/database/redis"
	"rockimserver/apis/rockim/service/group/v1/types"
)

type GroupMemberData struct {
	cache    cache.ValueCache[types.GroupMember]
	setCache cache.SortedSetCache[string]
}

func NewGroupMemberData(rds *redis.Client, cfg *cache.Config) *GroupMemberData {
	return &GroupMemberData{
		cache:    newValueCache[types.GroupMember](rds, cfg.Category(keyGroupMemberInfo)),
		setCache: newSortedSetCache[string](rds, cfg.Category(keyGroupMemberSet)),
	}
}

func (d *GroupMemberData) Add(ctx context.Context, member *types.GroupMember) (err error) {
	err = d.cache.Set(ctx, cache.Parts(member.ProductId, member.GroupId, member.Uid), member)
	if err != nil {
		return
	}
	setItem := &member.Uid
	return d.setCache.Add(ctx, cache.Parts(member.ProductId, member.GroupId), cache.Member[string](member.CreateTime, setItem))
}

func (d *GroupMemberData) Delete(ctx context.Context, productId string, groupId string, uid string) (err error) {
	err = d.cache.Delete(ctx, cache.Parts(productId, groupId, uid))
	if err != nil {
		return
	}
	setItem := &uid
	return d.setCache.DeleteMember(ctx, cache.Parts(productId, groupId), setItem)
}

func (d *GroupMemberData) Exists(ctx context.Context, productId string, groupId string, uid string) (exists bool, err error) {
	return d.cache.Exists(ctx, cache.Parts(productId, groupId, uid))
}
