package biz

import (
	"context"
	"github.com/soukengo/gopkg/component/idgen"
	"github.com/soukengo/gopkg/component/lock"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/group/biz/consts"
	"rockimserver/app/logic/group/biz/options"
	"strings"
	"time"
)

type ChatRoomUseCase struct {
	repo       GroupRepo
	memberRepo ChatRoomMemberRepo
	idgen      idgen.Generator
	lockBdr    lock.Builder
	memberMgr  *ChatRoomMemberManager
}

func NewChatRoomUseCase(repo GroupRepo, memberRepo ChatRoomMemberRepo, idgen idgen.Generator, lockBdr lock.Builder, memberMgr *ChatRoomMemberManager) *ChatRoomUseCase {
	return &ChatRoomUseCase{repo: repo, memberRepo: memberRepo, idgen: idgen, lockBdr: lockBdr, memberMgr: memberMgr}
}

func (uc *ChatRoomUseCase) FindByBizId(ctx context.Context, productId string, bizId string) (out string, err error) {
	return uc.repo.FindGroupId(ctx, productId, bizId)
}

func (uc *ChatRoomUseCase) FindById(ctx context.Context, productId string, groupId string) (out *types.Group, err error) {
	return uc.repo.FindById(ctx, productId, groupId)
}

func (uc *ChatRoomUseCase) Create(ctx context.Context, opts *options.ChatRoomCreateOptions) (out *types.Group, err error) {
	productId := opts.ProductId
	bizId := opts.BizId
	distributedLock := uc.lockBdr.Build(consts.LockKeyChatRoomCreate, productId, bizId)
	locked := distributedLock.TryLock(ctx)
	defer distributedLock.UnLock()
	if !locked {
		return
	}
	if len(bizId) > 0 {
		if strings.HasPrefix(bizId, consts.GroupBizIdPrefix) {
			err = ErrGroupBizIdInvalid
			return
		}
	} else {
		bizId, err = uc.idgen.GenID()
		if err != nil {
			return
		}
		bizId = consts.GroupBizIdPrefix + bizId
	}
	oldGroupId, err := uc.repo.FindGroupId(ctx, opts.ProductId, bizId)
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
		Id:         groupId,
		CreateTime: time.Now().UnixMilli(),
		UpdateTime: time.Now().UnixMilli(),
		Category:   enums.Group_CHATROOM,
		ProductId:  productId,
		BizId:      bizId,
		Name:       opts.Name,
		IconUrl:    opts.IconUrl,
		Fields:     opts.Fields,
		Owner:      opts.Owner,
	}
	err = uc.repo.Create(ctx, group)
	if err != nil {
		return
	}
	// 添加群主
	if len(opts.Owner) > 0 {
		err = uc.memberMgr.addMember(ctx, group, &options.ChatRoomMemberAddItem{Uid: opts.Owner, Role: enums.Group_OWNER})
		if err != nil {
			return
		}
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
