package cache

import (
	"context"
	"github.com/soukengo/gopkg/component/cache"
	"github.com/soukengo/gopkg/component/paginate"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared"
)

type ChatRoomMemberData struct {
	cache         cache.HashCache[types.GroupMember]
	setCache      cache.SortedSetCache[string]
	groupSetCache cache.SortedSetCache[string]
}

func NewChatRoomMemberData(mgr *cache.Manager, cfg *cache.Config) *ChatRoomMemberData {
	return &ChatRoomMemberData{
		cache:         cache.NewHashCache[types.GroupMember](mgr, cfg.Category(keyChatRoomMember)),
		setCache:      cache.NewSortedSetCache[string](mgr, cfg.Category(keyChatRoomMemberSet)),
		groupSetCache: cache.NewSortedSetCache[string](mgr, cfg.Category(keyUserChatRoomSet)),
	}
}

func (d *ChatRoomMemberData) Save(ctx context.Context, member *types.GroupMember) (err error) {
	err = d.cache.Put(ctx, cache.Parts(member.ProductId, member.GroupId), member.Uid, member)
	if err != nil {
		return
	}
	err = d.setCache.Add(ctx, cache.Parts(member.ProductId, member.GroupId), cache.Member[string](member.CreateTime, &member.Uid))
	if err != nil {
		return
	}
	err = d.groupSetCache.Add(ctx, cache.Parts(member.ProductId, member.Uid), cache.Member[string](member.CreateTime, &member.GroupId))
	return
}

func (d *ChatRoomMemberData) Delete(ctx context.Context, productId string, groupId string, uid string) (err error) {
	err = d.cache.DeleteField(ctx, cache.Parts(productId, groupId), uid)
	if err != nil {
		return
	}
	err = d.setCache.DeleteMember(ctx, cache.Parts(productId, groupId), &uid)
	if err != nil {
		return
	}
	return d.groupSetCache.DeleteMember(ctx, cache.Parts(productId, uid), &groupId)
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

func (d *ChatRoomMemberData) PaginateUid(ctx context.Context, productId string, groupId string, paging *shared.Paginating) (members []string, paginated *shared.Paginated, err error) {
	uids, p, err := d.setCache.Paginate(ctx, cache.Parts(productId, groupId), &paginate.Paginating{PageSize: paging.PageSize, PageNo: paging.PageNo})
	if errors.IsNotFound(err) {
		err = nil
	}
	paginated = &shared.Paginated{Total: p.Total}
	for _, uid := range uids {
		members = append(members, *uid)
	}
	return
}

func (d *ChatRoomMemberData) ListGroupIdByUid(ctx context.Context, productId string, uid string) (out []string, err error) {
	list, _, err := d.groupSetCache.Scan(ctx, cache.Parts(productId, uid), 0, 1000)
	if err != nil {
		return
	}
	// 排除已解散的群
	for _, item := range list {
		groupId := *item
		exists, err1 := d.setCache.ExistsMember(ctx, cache.Parts(productId, groupId), item)
		if err1 != nil && !cache.IsErrNoCache(err1) {
			err = err1
			return
		}
		if !exists {
			_ = d.groupSetCache.DeleteMember(ctx, cache.Parts(productId, uid), item)
			continue
		}
		out = append(out, groupId)
	}
	return
}
