package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

type GroupRepo interface {
	Find(ctx context.Context, productId string, customGroupId string) (*types.Group, error)
	IsMember(ctx context.Context, productId string, category enums.Group_Category, groupId string, uid string) (isMember bool, err error)
}
