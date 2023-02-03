package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/admin/manager/v1"
	"rockimserver/apis/rockim/api/admin/manager/v1/types"
	"rockimserver/app/access/admin/module/manager/biz"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"rockimserver/app/access/admin/module/manager/service/converter"
	"rockimserver/pkg/util/copier"
)

type SysTenantResourceService struct {
	uc *biz.SysTenantResourceUseCase
}

func NewSysTenantResourceService(uc *biz.SysTenantResourceUseCase) *SysTenantResourceService {
	return &SysTenantResourceService{uc: uc}
}

func (s *SysTenantResourceService) Create(ctx context.Context, in *v1.TenantResourceCreateRequest) (reply *v1.TenantResourceCreateResponse, err error) {
	req := &options.SysTenantResourceCreateOptions{}
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

func (s *SysTenantResourceService) Update(ctx context.Context, in *v1.TenantResourceUpdateRequest) (reply *v1.TenantResourceUpdateResponse, err error) {
	req := &options.SysTenantResourceUpdateOptions{}
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

func (s *SysTenantResourceService) Delete(ctx context.Context, in *v1.TenantResourceDeleteRequest) (reply *v1.TenantResourceDeleteResponse, err error) {
	err = s.uc.Delete(ctx, in.Id)
	if err != nil {
		return
	}
	reply = &v1.TenantResourceDeleteResponse{}
	return
}

func (s *SysTenantResourceService) ListTree(ctx context.Context, in *v1.TenantResourceTreeRequest) (reply *v1.TenantResourceTreeResponse, err error) {
	list, err := s.uc.List(ctx)
	if err != nil {
		return
	}
	resources := make([]*types.SysTenantResource, len(list))
	for i, item := range list {
		resources[i] = converter.ToManagerTenantResource(item)
	}
	return &v1.TenantResourceTreeResponse{List: converter.BuildTenantResourceTree(resources)}, nil
}
