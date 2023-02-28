package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/logic/group/biz/consts"
	"rockimserver/app/logic/group/biz/options"
	"rockimserver/pkg/component/lock"
	"rockimserver/pkg/errors"
	"time"
)

var (
	ErrChatRoomNotFound   = errors.NotFound(reasons.ChatRoom_CHATROOM_NOT_FOUND.String(), "chatroom not found")
	ErrChatRoomDuplicated = errors.Conflict(reasons.ChatRoom_CHATROOM_DUPLICATED.String(), "chatroom already exists")
)

type ChatRoomRepo interface {
	FindByCustomGroupId(ctx context.Context, productId string, customGroupId string) (string, error)
	FindById(ctx context.Context, productId string, groupId string) (*types.Group, error)
	Create(ctx context.Context, group *types.Group) error
	Delete(ctx context.Context, productId string, groupId string) error
}

type GroupIDRepo interface {
	GenerateID(ctx context.Context, productId string) (string, error)
	GenerateCustomGroupID(ctx context.Context, productId string) (string, error)
}

type ChatRoomUseCase struct {
	repo    ChatRoomRepo
	idRepo  GroupIDRepo
	lockBdr lock.Builder
}

func NewChatRoomUseCase(repo ChatRoomRepo, idRepo GroupIDRepo, lockBdr lock.Builder) *ChatRoomUseCase {
	return &ChatRoomUseCase{repo: repo, idRepo: idRepo, lockBdr: lockBdr}
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
		customGroupId, err = uc.idRepo.GenerateCustomGroupID(ctx, productId)
		if err != nil {
			return
		}
	}
	oldGroupId, err := uc.repo.FindByCustomGroupId(ctx, opts.ProductId, customGroupId)
	// 不是资源不存在的错误
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	if len(oldGroupId) > 0 {
		err = ErrChatRoomDuplicated
		return
	}
	groupId, err := uc.idRepo.GenerateCustomGroupID(ctx, productId)
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
	return
}
