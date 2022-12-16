package service

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/app/logic/platform/biz"
	"rockim/pkg/util/copier"
)

type PlatRoleService struct {
	v1.UnimplementedPlatRoleAPIServer
	uc *biz.PlatRoleUseCase
}

// NewPlatRoleService new a plat user service.
func NewPlatRoleService(uc *biz.PlatRoleUseCase) *PlatRoleService {
	return &PlatRoleService{uc: uc}
}

func (s *PlatRoleService) Create(ctx context.Context, in *v1.PlatRoleCreateRequest) (reply *v1.PlatRoleCreateResponse, err error) {
	req := &biz.PlatRoleCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatRoleCreateResponse{}
	return
}

func (s *PlatRoleService) Update(ctx context.Context, in *v1.PlatRoleUpdateRequest) (reply *v1.PlatRoleUpdateResponse, err error) {
	req := &biz.PlatRoleUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatRoleUpdateResponse{}
	return
}

func (s *PlatRoleService) Delete(ctx context.Context, in *v1.PlatRoleDeleteRequest) (reply *v1.PlatRoleDeleteResponse, err error) {
	req := &biz.PlatRoleDeleteRequest{Id: in.Id}
	err = s.uc.Delete(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatRoleDeleteResponse{}
	return
}

func (s *PlatRoleService) Paging(ctx context.Context, in *v1.PlatRolePagingRequest) (reply *v1.PlatRolePagingResponse, err error) {
	req := &biz.PlatRolePagingRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	res, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.PlatRolePagingResponse{List: res.List, Paginate: res.Paginate}
	return
}

func (s *PlatRoleService) ListResourceId(ctx context.Context, in *v1.PlatRoleResourceIdListRequest) (*v1.PlatRoleResourceIdListResponse, error) {
	list, err := s.uc.ListResourceId(ctx, in.RoleIds)
	if err != nil {
		return nil, err
	}
	return &v1.PlatRoleResourceIdListResponse{List: list}, nil
}
