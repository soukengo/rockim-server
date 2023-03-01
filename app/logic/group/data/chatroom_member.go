package data

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/data/cache"
	"rockimserver/pkg/errors"
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
func (r *chatRoomMemberRepo) DeleteAll(ctx context.Context, productId string, groupId string) error {
	return r.cache.DeleteAll(ctx, productId, groupId)
}

func (r *chatRoomMemberRepo) Exists(ctx context.Context, productId string, groupId string, uid string) (exists bool, err error) {
	exists, err = r.cache.Exists(ctx, productId, groupId, uid)
	if errors.IsNotFound(err) {
		err = nil
	}
	return
}

func (r *chatRoomMemberRepo) Find(ctx context.Context, productId string, groupId string, uid string) (ret *types.GroupMember, err error) {
	ret, err = r.cache.Find(ctx, productId, groupId, uid)
	if errors.IsNotFound(err) {
		err = biz.ErrGroupMemberNotFound
	}
	return
}

func (r *chatRoomMemberRepo) List(ctx context.Context, productId string, groupId string, uids []string) (ret []*types.GroupMember, err error) {
	ret, err = r.cache.List(ctx, productId, groupId, uids)
	if errors.IsNotFound(err) {
		err = nil
	}
	return
}

func (r *chatRoomMemberRepo) PaginateUid(ctx context.Context, productId string, groupId string, paginate *shared.Paginating) (ret []string, paginated *shared.Paginated, err error) {
	ret, paginated, err = r.cache.PaginateUid(ctx, productId, groupId, paginate)
	if errors.IsNotFound(err) {
		err = nil
	}
	return
}
