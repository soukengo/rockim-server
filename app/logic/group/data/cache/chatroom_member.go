package cache

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared"
	"rockimserver/pkg/component/cache"
	"rockimserver/pkg/component/database/redis"
	"rockimserver/pkg/errors"
)

type ChatRoomMemberData struct {
	cache    cache.HashCache[types.GroupMember]
	setCache cache.SortedSetCache[string]
}

func NewChatRoomMemberData(rds *redis.Client, cfg *cache.Config) *ChatRoomMemberData {
	return &ChatRoomMemberData{
		cache:    newHashCache[types.GroupMember](rds, cfg.Category(keyChatRoomMemberHash)),
		setCache: newSortedSetCache[string](rds, cfg.Category(keyChatRoomMemberSet)),
	}
}

func (d *ChatRoomMemberData) Save(ctx context.Context, member *types.GroupMember) (err error) {
	err = d.cache.Put(ctx, cache.Parts(member.ProductId, member.GroupId), member.Uid, member)
	if err != nil {
		return
	}
	setItem := &member.Uid
	return d.setCache.Add(ctx, cache.Parts(member.ProductId, member.GroupId), cache.Member[string](member.CreateTime, setItem))
}

func (d *ChatRoomMemberData) Delete(ctx context.Context, productId string, groupId string, uid string) (err error) {
	err = d.cache.DeleteField(ctx, cache.Parts(productId, groupId), uid)
	if err != nil {
		return
	}
	setItem := &uid
	return d.setCache.DeleteMember(ctx, cache.Parts(productId, groupId), setItem)
}

func (d *ChatRoomMemberData) DeleteAll(ctx context.Context, productId string, groupId string) (err error) {
	err = d.cache.Delete(ctx, cache.Parts(productId, groupId))
	if err != nil {
		return
	}
	return d.setCache.Delete(ctx, cache.Parts(productId, groupId))
}

func (d *ChatRoomMemberData) Exists(ctx context.Context, productId string, groupId string, uid string) (exists bool, err error) {
	return d.cache.Exists(ctx, cache.Parts(productId, groupId, uid))
}

func (d *ChatRoomMemberData) Find(ctx context.Context, productId string, groupId string, uid string) (member *types.GroupMember, err error) {
	return d.cache.Get(ctx, cache.Parts(productId, groupId), uid)
}
func (d *ChatRoomMemberData) List(ctx context.Context, productId string, groupId string, uids []string) (members []*types.GroupMember, err error) {
	return d.cache.MGet(ctx, cache.Parts(productId, groupId), uids)
}
func (d *ChatRoomMemberData) PaginateUid(ctx context.Context, productId string, groupId string, paginate *shared.Paginating) (members []string, paginated *shared.Paginated, err error) {
	uids, paginated, err := d.setCache.Paginate(ctx, cache.Parts(productId, groupId), paginate)
	if errors.IsNotFound(err) {
		err = nil
	}
	for _, uid := range uids {
		members = append(members, *uid)
	}
	return
}
