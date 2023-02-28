package data

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/data/cache"
)

type groupMemberRepo struct {
	cache *cache.GroupMemberData
}

func NewGroupMemberRepo(cache *cache.GroupMemberData) biz.GroupMemberRepo {
	return &groupMemberRepo{cache: cache}
}

func (r *groupMemberRepo) Add(ctx context.Context, member *types.GroupMember) error {
	return r.cache.Add(ctx, member)
}

func (r *groupMemberRepo) Delete(ctx context.Context, productId string, groupId string, uid string) error {
	return r.cache.Delete(ctx, productId, groupId, uid)
}

func (r *groupMemberRepo) Exists(ctx context.Context, productId string, groupId string, uid string) (exists bool, err error) {
	return r.cache.Exists(ctx, productId, groupId, uid)
}
