package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/group/biz/options"
)

type GroupMemberRepo interface {
	Add(ctx context.Context, member *types.GroupMember) error
	Delete(ctx context.Context, productId string, groupId string, uid string) error
	Exists(ctx context.Context, productId string, groupId string, uid string) (bool, error)
	Find(ctx context.Context, productId string, groupId string, uid string) (*types.GroupMember, error)
	List(ctx context.Context, productId string, groupId string, uids []string) ([]*types.GroupMember, error)
}

type GroupMemberUseCase struct {
	chatRoomRepo    ChatRoomMemberRepo
	groupMemberRepo GroupMemberRepo
}

func NewGroupMemberUseCase(chatRoomRepo ChatRoomMemberRepo, groupMemberRepo GroupMemberRepo) *GroupMemberUseCase {
	return &GroupMemberUseCase{chatRoomRepo: chatRoomRepo, groupMemberRepo: groupMemberRepo}
}

func (uc *GroupMemberUseCase) IsMember(ctx context.Context, opts *options.GroupMemberCheckOptions) (isMember bool, err error) {
	if opts.Category == enums.Group_CHATROOM {
		return uc.chatRoomRepo.Exists(ctx, opts.ProductId, opts.GroupId, opts.Uid)
	}
	return uc.groupMemberRepo.Exists(ctx, opts.ProductId, opts.GroupId, opts.Uid)
}
func (uc *GroupMemberUseCase) Find(ctx context.Context, opts *options.GroupMemberFindOptions) (member *types.GroupMember, err error) {
	if opts.Category == enums.Group_CHATROOM {
		return uc.chatRoomRepo.Find(ctx, opts.ProductId, opts.GroupId, opts.Uid)
	}
	return uc.groupMemberRepo.Find(ctx, opts.ProductId, opts.GroupId, opts.Uid)
}

func (uc *GroupMemberUseCase) List(ctx context.Context, opts *options.GroupMemberListOptions) (members []*types.GroupMember, err error) {
	if opts.Category == enums.Group_CHATROOM {
		return uc.chatRoomRepo.List(ctx, opts.ProductId, opts.GroupId, opts.Uids)
	}
	return uc.groupMemberRepo.List(ctx, opts.ProductId, opts.GroupId, opts.Uids)
}

func (uc *GroupMemberUseCase) PaginateUid(ctx context.Context, opts *options.GroupMemberIdPaginateOptions) (members []string, paginated *shared.Paginated, err error) {
	return uc.chatRoomRepo.PaginateUid(ctx, opts.ProductId, opts.GroupId, opts.Paginate)
}

func (uc *GroupMemberUseCase) ListGroupIdByUid(ctx context.Context, opts *options.GroupIdListByUidOptions) (members []string, err error) {
	return uc.chatRoomRepo.ListGroupIdByUid(ctx, opts.ProductId, opts.Uid)
}
