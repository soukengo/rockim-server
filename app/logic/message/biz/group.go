package biz

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
)

type GroupRepo interface {
	Find(ctx context.Context, productId string, bizId string) (*types.Group, error)
}
