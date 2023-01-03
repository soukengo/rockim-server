package service

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/app/logic/platform/biz"
	"rockim/pkg/util/copier"
)

type TenantResourceService struct {
	v1.UnimplementedTenantResourceAPIServer
	uc *biz.TenantResourceUseCase
}

// NewTenantResourceService new a plat resource service.
func NewTenantResourceService(uc *biz.TenantResourceUseCase) *TenantResourceService {
	return &TenantResourceService{uc: uc}
}

func (s *TenantResourceService) Create(ctx context.Context, in *v1.TenantResourceCreateRequest) (reply *v1.TenantResourceCreateResponse, err error) {
	req := &biz.TenantResourceCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.TenantResourceCreateResponse{}
	return
}

func (s *TenantResourceService) Update(ctx context.Context, in *v1.TenantResourceUpdateRequest) (reply *v1.TenantResourceUpdateResponse, err error) {
	req := &biz.TenantResourceUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.TenantResourceUpdateResponse{}
	return
}

func (s *TenantResourceService) Delete(ctx context.Context, in *v1.TenantResourceDeleteRequest) (reply *v1.TenantResourceDeleteResponse, err error) {
	req := &biz.TenantResourceDeleteRequest{Id: in.Id}
	err = s.uc.Delete(ctx, req)
	if err != nil {
		return
	}
	reply = &v1.TenantResourceDeleteResponse{}
	return
}

func (s *TenantResourceService) List(ctx context.Context, request *v1.TenantResourceListRequest) (reply *v1.TenantResourceListResponse, err error) {
	list, err := s.uc.List(ctx)
	if err != nil {
		return
	}
	return &v1.TenantResourceListResponse{List: list}, nil
}

func (s *TenantResourceService) ListByIds(ctx context.Context, request *v1.TenantResourceListByIdsRequest) (reply *v1.TenantResourceListByIdsResponse, err error) {
	list, err := s.uc.ListByIds(ctx, request.Ids)
	if err != nil {
		return
	}
	reply = &v1.TenantResourceListByIdsResponse{List: list}
	return
}
