package manager

import (
	"context"
	adminV1 "rockim/api/rockim/admin/manager/v1"
	adminTypes "rockim/api/rockim/admin/manager/v1/types"
	"rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
	"rockim/pkg/util/copier"
)

type PlatRoleService struct {
	uc *manager.PlatRoleUseCase
}

func NewPlatRoleService(uc *manager.PlatRoleUseCase) *PlatRoleService {
	return &PlatRoleService{uc: uc}
}

func (s *PlatRoleService) Create(ctx context.Context, in *adminV1.PlatRoleCreateRequest) (reply *adminV1.PlatRoleCreateResponse, err error) {
	req := &v1.PlatRoleCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatRoleCreateResponse{}
	return
}

func (s *PlatRoleService) Update(ctx context.Context, in *adminV1.PlatRoleUpdateRequest) (reply *adminV1.PlatRoleUpdateResponse, err error) {
	req := &v1.PlatRoleUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatRoleUpdateResponse{}
	return
}

func (s *PlatRoleService) Delete(ctx context.Context, in *adminV1.PlatRoleDeleteRequest) (reply *adminV1.PlatRoleDeleteResponse, err error) {
	req := &v1.PlatRoleDeleteRequest{Id: in.Id}
	err = s.uc.Delete(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatRoleDeleteResponse{}
	return
}

func (s *PlatRoleService) Paging(ctx context.Context, in *adminV1.PlatRolePagingRequest) (reply *adminV1.PlatRolePagingResponse, err error) {
	req := &v1.PlatRolePagingRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	ret, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	list := make([]*adminTypes.PlatRole, len(ret.List))
	for i, item := range ret.List {
		list[i] = convertRole(item)
	}
	return &adminV1.PlatRolePagingResponse{List: list, Paginate: ret.Paginate}, nil
}

func (s *PlatRoleService) ListResourceId(ctx context.Context, in *adminV1.PlatRoleResourceIdListRequest) (reply *adminV1.PlatRoleResourceIdListResponse, err error) {
	list, err := s.uc.ListResourceId(ctx, in.RoleId)
	if err != nil {
		return
	}
	return &adminV1.PlatRoleResourceIdListResponse{List: list}, nil
}
func (s *PlatRoleService) SaveResourceId(ctx context.Context, in *adminV1.PlatRoleResourceIdSaveRequest) (reply *adminV1.PlatRoleResourceIdSaveResponse, err error) {
	err = s.uc.SaveResourceId(ctx, in.RoleId, in.ResourceIds)
	if err != nil {
		return
	}
	return &adminV1.PlatRoleResourceIdSaveResponse{}, nil
}

func convertRole(source *types.PlatRole) *adminTypes.PlatRole {
	return &adminTypes.PlatRole{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		Name:       source.Name,
	}
}
