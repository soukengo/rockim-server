package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/pkg/errors"
)

var (
	ErrGroupNotFound   = errors.NotFound(reasons.Group_GROUP_NOT_FOUND.String(), "group not found")
	ErrGroupDuplicated = errors.Conflict(reasons.Group_GROUP_DUPLICATED.String(), "group already exists")
)

type GroupRepo interface {
	FindByCustomGroupId(ctx context.Context, productId string, customGroupId string) (string, error)
	FindById(ctx context.Context, productId string, groupId string) (*types.Group, error)
	Create(ctx context.Context, group *types.Group) error
	Delete(ctx context.Context, group *types.Group) error
}

type GroupIDRepo interface {
	GenerateID(ctx context.Context, productId string) (string, error)
	GenerateCustomGroupID(ctx context.Context, productId string) (string, error)
}

type GroupMemberRepo interface {
	Add(ctx context.Context, member *types.GroupMember) error
	Delete(ctx context.Context, productId string, groupId string, uid string) error
	Exists(ctx context.Context, productId string, groupId string, uid string) (bool, error)
}
