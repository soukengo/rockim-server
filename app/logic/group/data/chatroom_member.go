package data

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/data/cache"
)

type chatRoomMemberRepo struct {
	cache *cache.ChatRoomMemberData
}

func NewChatRoomMemberRepo(cache *cache.ChatRoomMemberData) biz.ChatRoomMemberRepo {
	return &chatRoomMemberRepo{cache: cache}
}

func (r *chatRoomMemberRepo) Add(ctx context.Context, member *types.GroupMember) error {
	return r.cache.Save(ctx, member)
}
func (r *chatRoomMemberRepo) Update(ctx context.Context, member *types.GroupMember) error {
	return r.cache.Save(ctx, member)
}

func (r *chatRoomMemberRepo) Delete(ctx context.Context, productId string, groupId string, uid string) error {
	return r.cache.Delete(ctx, productId, groupId, uid)
}

func (r *chatRoomMemberRepo) Exists(ctx context.Context, productId string, groupId string, uid string) (exists bool, err error) {
	return r.cache.Exists(ctx, productId, groupId, uid)
}
