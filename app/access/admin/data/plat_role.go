package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/app/access/admin/biz/manager"
)

type platRoleRepo struct {
	uac v1.PlatRoleAPIClient
}

func NewPlatRoleRepo(uac v1.PlatRoleAPIClient) manager.PlatRoleRepo {
	return &platRoleRepo{uac: uac}
}

func (r *platRoleRepo) ListResourceId(ctx context.Context, roleIds []string) ([]string, error) {
	ret, err := r.uac.ListResourceId(ctx, &v1.PlatRoleResourceIdListRequest{
		RoleIds: roleIds,
	})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}
