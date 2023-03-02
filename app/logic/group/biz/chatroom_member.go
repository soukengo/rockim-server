package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/group/biz/consts"
	"rockimserver/app/logic/group/biz/options"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
	"time"
)

type ChatRoomMemberRepo interface {
	Add(ctx context.Context, member *types.GroupMember) error
	Update(ctx context.Context, member *types.GroupMember) error
	DeleteAll(ctx context.Context, productId string, groupId string) error
	Delete(ctx context.Context, productId string, groupId string, uid string) error
	Exists(ctx context.Context, productId string, groupId string, uid string) (bool, error)
	Find(ctx context.Context, productId string, groupId string, uid string) (*types.GroupMember, error)
	List(ctx context.Context, productId string, groupId string, uids []string) ([]*types.GroupMember, error)
	PaginateUid(ctx context.Context, productId string, groupId string, paginate *shared.Paginating) ([]string, *shared.Paginated, error)
}

type ChatRoomMemberUseCase struct {
	groupRepo GroupRepo
	repo      ChatRoomMemberRepo
	lockBdr   lock.Builder
	idgen     idgen.Generator
	memberMgr *ChatRoomMemberManager
}

func NewChatRoomMemberUseCase(groupRepo GroupRepo, repo ChatRoomMemberRepo, lockBdr lock.Builder, idgen idgen.Generator, memberMgr *ChatRoomMemberManager) *ChatRoomMemberUseCase {
	return &ChatRoomMemberUseCase{groupRepo: groupRepo, repo: repo, lockBdr: lockBdr, idgen: idgen, memberMgr: memberMgr}
}

func (uc *ChatRoomMemberUseCase) Join(ctx context.Context, opts *options.ChatRoomJoinOptions) (err error) {
	productId := opts.ProductId
	groupId := opts.GroupId
	uid := opts.Uid
	distributedLock := uc.lockBdr.Build(consts.LockKeyChatRoomJoin, productId, groupId, uid)
	locked := distributedLock.TryLock(ctx)
	defer distributedLock.UnLock()
	if !locked {
		return
	}
	group, err := uc.groupRepo.FindById(ctx, productId, groupId)
	if err != nil {
		return
	}
	isMember, err := uc.repo.Exists(ctx, productId, groupId, uid)
	if err != nil {
		return
	}
	if isMember {
		return
	}
	err = uc.memberMgr.addMember(ctx, group, &options.ChatRoomMemberAddItem{Uid: uid, Role: enums.Group_ORDINARY})
	return
}
func (uc *ChatRoomMemberUseCase) Quit(ctx context.Context, opts *options.ChatRoomQuitOptions) (err error) {
	productId := opts.ProductId
	groupId := opts.GroupId
	uid := opts.Uid
	distributedLock := uc.lockBdr.Build(consts.LockKeyChatRoomQuit, productId, groupId, uid)
	locked := distributedLock.TryLock(ctx)
	defer distributedLock.UnLock()
	if !locked {
		return
	}
	isMember, err := uc.repo.Exists(ctx, productId, groupId, uid)
	if err != nil {
		return
	}
	if isMember {
		return
	}
	err = uc.repo.Delete(ctx, productId, groupId, uid)
	if err != nil {
		return
	}
	// todo send notification
	return
}
func (uc *ChatRoomMemberUseCase) IsMember(ctx context.Context, opts *options.ChatRoomMemberCheckOptions) (isMember bool, err error) {
	return uc.repo.Exists(ctx, opts.ProductId, opts.GroupId, opts.Uid)
}
func (uc *ChatRoomMemberUseCase) Find(ctx context.Context, opts *options.ChatRoomMemberFindOptions) (member *types.GroupMember, err error) {
	return uc.repo.Find(ctx, opts.ProductId, opts.GroupId, opts.Uid)
}

func (uc *ChatRoomMemberUseCase) List(ctx context.Context, opts *options.ChatRoomMemberListOptions) (members []*types.GroupMember, err error) {
	return uc.repo.List(ctx, opts.ProductId, opts.GroupId, opts.Uids)
}
func (uc *ChatRoomMemberUseCase) PaginateUid(ctx context.Context, opts *options.ChatRoomMemberIdPaginateOptions) (members []string, paginated *shared.Paginated, err error) {
	return uc.repo.PaginateUid(ctx, opts.ProductId, opts.GroupId, opts.Paginate)
}

type ChatRoomMemberManager struct {
	groupRepo GroupRepo
	repo      ChatRoomMemberRepo
	idgen     idgen.Generator
}

func NewChatRoomMemberManager(groupRepo GroupRepo, repo ChatRoomMemberRepo, idgen idgen.Generator) *ChatRoomMemberManager {
	return &ChatRoomMemberManager{groupRepo: groupRepo, repo: repo, idgen: idgen}
}

// addMember add chatroom member
func (mgr *ChatRoomMemberManager) addMember(ctx context.Context, group *types.Group, item *options.ChatRoomMemberAddItem) (err error) {
	productId := group.ProductId
	groupId := group.Id
	uid := item.Uid
	recordId, err := mgr.idgen.GenID()
	if err != nil {
		return
	}
	role := item.Role
	if role == enums.Group_OWNER && group.Owner != uid {
		role = enums.Group_ORDINARY
	}
	member := &types.GroupMember{
		Id:         recordId,
		CreateTime: time.Now().UnixMilli(),
		ProductId:  productId,
		GroupId:    groupId,
		Uid:        uid,
		Role:       role,
		Fields:     item.Fields,
	}
	err = mgr.repo.Add(ctx, member)
	if err != nil {
		return
	}
	// todo send notification
	return
}
