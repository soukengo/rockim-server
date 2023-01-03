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

type TenantService struct {
	uc *manager.TenantUseCase
}

func NewTenantService(uc *manager.TenantUseCase) *TenantService {
	return &TenantService{uc: uc}
}

func (s *TenantService) Create(ctx context.Context, in *adminV1.TenantCreateRequest) (reply *adminV1.TenantCreateResponse, err error) {
	req := &v1.TenantCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.TenantCreateResponse{}
	return
}

func (s *TenantService) Update(ctx context.Context, in *adminV1.TenantUpdateRequest) (reply *adminV1.TenantUpdateResponse, err error) {
	req := &v1.TenantUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.TenantUpdateResponse{}
	return
}

func (s *TenantService) Paging(ctx context.Context, in *adminV1.TenantPagingRequest) (reply *adminV1.TenantPagingResponse, err error) {
	req := &v1.TenantPagingRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	ret, err := s.uc.Paging(ctx, req)
	if err != nil {
		return
	}
	list := make([]*adminTypes.Tenant, len(ret.List))
	for i, item := range ret.List {
		list[i] = convertTenant(item)
	}
	return &adminV1.TenantPagingResponse{List: list, Paginate: ret.Paginate}, nil
}

func convertTenant(source *types.Tenant) *adminTypes.Tenant {
	return &adminTypes.Tenant{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		Name:       source.Name,
		Email:      source.Email,
		Status:     source.Status,
	}
}
