package service

import (
	"context"
	adminV1 "rockimserver/apis/rockim/api/admin/manager/v1"
	apiTypes "rockimserver/apis/rockim/api/admin/manager/v1/types"
	"rockimserver/app/access/admin/module/manager/biz"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"rockimserver/app/access/admin/module/manager/service/converter"
	"rockimserver/pkg/util/copier"
)

type SysRoleService struct {
	uc *biz.SysRoleUseCase
}

func NewSysRoleService(uc *biz.SysRoleUseCase) *SysRoleService {
	return &SysRoleService{uc: uc}
}

func (s *SysRoleService) Create(ctx context.Context, in *adminV1.SysRoleCreateRequest) (reply *adminV1.SysRoleCreateResponse, err error) {
	req := &options.SysRoleCreateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.SysRoleCreateResponse{}
	return
}

func (s *SysRoleService) Update(ctx context.Context, in *adminV1.SysRoleUpdateRequest) (reply *adminV1.SysRoleUpdateResponse, err error) {
	req := &options.SysRoleUpdateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.SysRoleUpdateResponse{}
	return
}

func (s *SysRoleService) Delete(ctx context.Context, in *adminV1.SysRoleDeleteRequest) (reply *adminV1.SysRoleDeleteResponse, err error) {
	err = s.uc.Delete(ctx, in.Id)
	if err != nil {
		return
	}
	reply = &adminV1.SysRoleDeleteResponse{}
	return
}

func (s *SysRoleService) Paging(ctx context.Context, in *adminV1.SysRolePagingRequest) (reply *adminV1.SysRolePagingResponse, err error) {
	req := &options.SysRolePagingOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	roles, p, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	list := make([]*apiTypes.SysRole, len(roles))
	for i, item := range roles {
		list[i] = converter.ToManagerSysRole(item)
	}
	return &adminV1.SysRolePagingResponse{List: list, Paginate: p}, nil
}

func (s *SysRoleService) ListResourceId(ctx context.Context, in *adminV1.SysRoleResourceIdListRequest) (reply *adminV1.SysRoleResourceIdListResponse, err error) {
	list, err := s.uc.ListResourceId(ctx, in.RoleId)
	if err != nil {
		return
	}
	return &adminV1.SysRoleResourceIdListResponse{List: list}, nil
}
func (s *SysRoleService) SaveResourceId(ctx context.Context, in *adminV1.SysRoleResourceIdSaveRequest) (reply *adminV1.SysRoleResourceIdSaveResponse, err error) {
	err = s.uc.SaveResourceId(ctx, in.RoleId, in.ResourceIds)
	if err != nil {
		return
	}
	return &adminV1.SysRoleResourceIdSaveResponse{}, nil
}
