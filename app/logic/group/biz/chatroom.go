package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/group/biz/consts"
	"rockimserver/app/logic/group/biz/options"
	"rockimserver/pkg/component/idgen"
	"rockimserver/pkg/component/lock"
	"rockimserver/pkg/errors"
	"time"
)

type ChatRoomUseCase struct {
	repo       GroupRepo
	memberRepo ChatRoomMemberRepo
	idgen      idgen.Generator
	lockBdr    lock.Builder
}

func NewChatRoomUseCase(repo GroupRepo, memberRepo ChatRoomMemberRepo, idgen idgen.Generator, lockBdr lock.Builder) *ChatRoomUseCase {
	return &ChatRoomUseCase{repo: repo, memberRepo: memberRepo, idgen: idgen, lockBdr: lockBdr}
}

func (uc *ChatRoomUseCase) FindByCustomGroupId(ctx context.Context, productId string, customGroupId string) (out string, err error) {
	return uc.repo.FindByCustomGroupId(ctx, productId, customGroupId)
}

func (uc *ChatRoomUseCase) FindById(ctx context.Context, productId string, groupId string) (out *types.Group, err error) {
	return uc.repo.FindById(ctx, productId, groupId)
}

func (uc *ChatRoomUseCase) Create(ctx context.Context, opts *options.ChatRoomCreateOptions) (out *types.Group, err error) {
	productId := opts.ProductId
	customGroupId := opts.CustomGroupId
	distributedLock := uc.lockBdr.Build(consts.LockKeyChatRoomCreate, productId, customGroupId)
	locked := distributedLock.TryLock(ctx)
	defer distributedLock.UnLock()
	if !locked {
		return
	}
	if len(customGroupId) == 0 {
		customGroupId, err = uc.idgen.GenID()
		if err != nil {
			return
		}
		customGroupId = consts.ChatRoomIDPrefix + customGroupId
	}
	oldGroupId, err := uc.repo.FindByCustomGroupId(ctx, opts.ProductId, customGroupId)
	// 不是资源不存在的错误
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	if len(oldGroupId) > 0 {
		err = ErrGroupDuplicated
		return
	}
	groupId, err := uc.idgen.GenID()
	if err != nil {
		return
	}
	group := &types.Group{
		Id:            groupId,
		CreateTime:    time.Now().UnixMilli(),
		UpdateTime:    time.Now().UnixMilli(),
		Category:      enums.Group_CHATROOM,
		ProductId:     productId,
		CustomGroupId: customGroupId,
		Name:          opts.Name,
		IconUrl:       opts.IconUrl,
		Fields:        opts.Fields,
	}
	err = uc.repo.Create(ctx, group)
	if err != nil {
		return
	}
	out = group
	return
}
func (uc *ChatRoomUseCase) Dismiss(ctx context.Context, opts *options.ChatRoomDismissOptions) (err error) {
	productId := opts.ProductId
	groupId := opts.GroupId
	distributedLock := uc.lockBdr.Build(consts.LockKeyChatRoomDismiss, productId, groupId)
	locked := distributedLock.TryLock(ctx)
	defer distributedLock.UnLock()
	if !locked {
		return
	}
	group, err := uc.repo.FindById(ctx, productId, groupId)
	if err != nil {
		return
	}
	// 移除所有成员
	err = uc.memberRepo.DeleteAll(ctx, productId, groupId)
	if err != nil {
		return
	}
	// 删除群信息
	err = uc.repo.Delete(ctx, group)
	if err != nil {
		return
	}
	// todo send notification
	return
}
