package data

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/app/access/admin/biz/manager"
)

type platRoleRepo struct {
	ac v1.PlatRoleAPIClient
}

func NewPlatRoleRepo(uac v1.PlatRoleAPIClient) manager.PlatRoleRepo {
	return &platRoleRepo{ac: uac}
}

func (r *platRoleRepo) Create(ctx context.Context, request *v1.PlatRoleCreateRequest) (err error) {
	_, err = r.ac.Create(ctx, request)
	return
}

func (r *platRoleRepo) Update(ctx context.Context, request *v1.PlatRoleUpdateRequest) (err error) {
	_, err = r.ac.Update(ctx, request)
	return
}

func (r *platRoleRepo) Delete(ctx context.Context, request *v1.PlatRoleDeleteRequest) (err error) {
	_, err = r.ac.Delete(ctx, request)
	return
}

func (r *platRoleRepo) Paging(ctx context.Context, request *v1.PlatRolePagingRequest) (reply *v1.PlatRolePagingResponse, err error) {
	return r.ac.Paging(ctx, request)
}

func (r *platRoleRepo) ListResourceId(ctx context.Context, roleIds []string) ([]string, error) {
	ret, err := r.ac.ListResourceId(ctx, &v1.PlatRoleResourceIdListRequest{
		RoleIds: roleIds,
	})
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}

func (r *platRoleRepo) SaveResourceId(ctx context.Context, roleId string, resourceIds []string) (err error) {
	_, err = r.ac.SaveResourceId(ctx, &v1.PlatRoleResourceIdSaveRequest{
		RoleId:      roleId,
		ResourceIds: resourceIds,
	})
	return
}
