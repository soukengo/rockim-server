package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/pkg/errors"
)

var (
	ErrGroupNotFound             = errors.NotFound(reasons.Group_GROUP_NOT_FOUND.String(), "group not found")
	ErrGroupDuplicated           = errors.Conflict(reasons.Group_GROUP_DUPLICATED.String(), "group already exists")
	ErrGroupMemberNotFound       = errors.Conflict(reasons.Group_GROUP_MEMBER_NOT_FOUND.String(), "group member not found")
	ErrGroupCustomGroupIdInvalid = errors.Conflict(reasons.Group_CUSTOM_GROUP_ID_INVALID.String(), "group member not found")
)

type GroupRepo interface {
	FindByCustomGroupId(ctx context.Context, productId string, customGroupId string) (string, error)
	FindById(ctx context.Context, productId string, groupId string) (*types.Group, error)
	Create(ctx context.Context, group *types.Group) error
	Delete(ctx context.Context, group *types.Group) error
}

type GroupMemberRepo interface {
	Add(ctx context.Context, member *types.GroupMember) error
	Delete(ctx context.Context, productId string, groupId string, uid string) error
	Exists(ctx context.Context, productId string, groupId string, uid string) (bool, error)
}
