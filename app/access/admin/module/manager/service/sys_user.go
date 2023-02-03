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

type SysUserService struct {
	uc *biz.SysUserUseCase
}

func NewSysUserService(uc *biz.SysUserUseCase) *SysUserService {
	return &SysUserService{uc: uc}
}

func (s *SysUserService) Create(ctx context.Context, in *adminV1.SysUserCreateRequest) (reply *adminV1.SysUserCreateResponse, err error) {
	req := &options.SysUserCreateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.SysUserCreateResponse{}
	return
}

func (s *SysUserService) Update(ctx context.Context, in *adminV1.SysUserUpdateRequest) (reply *adminV1.SysUserUpdateResponse, err error) {
	req := &options.SysUserUpdateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.SysUserUpdateResponse{}
	return
}

func (s *SysUserService) Delete(ctx context.Context, in *adminV1.SysUserDeleteRequest) (reply *adminV1.SysUserDeleteResponse, err error) {
	err = s.uc.Delete(ctx, in.Id)
	if err != nil {
		return
	}
	reply = &adminV1.SysUserDeleteResponse{}
	return
}

func (s *SysUserService) Paging(ctx context.Context, in *adminV1.SysUserPagingRequest) (reply *adminV1.SysUserPagingResponse, err error) {
	req := &options.SysUserPagingOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	users, p, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	list := make([]*apiTypes.SysUser, len(users))
	for i, item := range users {
		list[i] = converter.ToManagerSysUser(item)
	}
	return &adminV1.SysUserPagingResponse{List: list, Paginate: p}, nil
}

func (s *SysUserService) ListRoleId(ctx context.Context, in *adminV1.SysUserRoleIdListRequest) (reply *adminV1.SysUserRoleIdListResponse, err error) {
	list, err := s.uc.ListRoleId(ctx, in.UserId)
	if err != nil {
		return
	}
	return &adminV1.SysUserRoleIdListResponse{List: list}, nil
}
func (s *SysUserService) SaveRoleId(ctx context.Context, in *adminV1.SysUserRoleIdSaveRequest) (reply *adminV1.SysUserRoleIdSaveResponse, err error) {
	err = s.uc.SaveRoleId(ctx, in.UserId, in.RoleIds)
	if err != nil {
		return
	}
	return &adminV1.SysUserRoleIdSaveResponse{}, nil
}
