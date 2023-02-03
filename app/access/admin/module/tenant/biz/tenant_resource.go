package biz

import (
	"context"
	"rockimserver/app/access/admin/module/tenant/biz/types"
)

type SysTenantResourceRepo interface {
	List(context.Context) ([]*types.SysTenantResource, error)
}
