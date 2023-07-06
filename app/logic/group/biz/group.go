package biz

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/reasons"
)

var (
	ErrGroupNotFound             = errors.NotFound(reasons.Group_GROUP_NOT_FOUND.String(), "group not found")
	ErrGroupDuplicated           = errors.Conflict(reasons.Group_GROUP_DUPLICATED.String(), "group already exists")
	ErrGroupMemberNotFound       = errors.Conflict(reasons.Group_GROUP_MEMBER_NOT_FOUND.String(), "group member not found")
	ErrGroupCustomGroupIdInvalid = errors.Conflict(reasons.Group_CUSTOM_GROUP_ID_INVALID.String(), "group member not found")
)

type GroupRepo interface {
	FindGroupId(ctx context.Context, productId string, customGroupId string) (string, error)
	FindById(ctx context.Context, productId string, groupId string) (*types.Group, error)
	Create(ctx context.Context, group *types.Group) error
	Delete(ctx context.Context, group *types.Group) error
}

type GroupUseCase struct {
	repo GroupRepo
}

func NewGroupUseCase(repo GroupRepo) *GroupUseCase {
	return &GroupUseCase{repo: repo}
}

func (uc *GroupUseCase) FindGroupId(ctx context.Context, productId string, customGroupId string) (string, error) {
	return uc.repo.FindGroupId(ctx, productId, customGroupId)
}
func (uc *GroupUseCase) FindById(ctx context.Context, productId string, groupId string) (*types.Group, error) {
	return uc.repo.FindById(ctx, productId, groupId)
}
