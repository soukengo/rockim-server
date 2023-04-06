package service

import (
	"context"
	"github.com/soukengo/gopkg/util/copier"
	v1 "rockimserver/apis/rockim/api/admin/tenant/v1"
	"rockimserver/apis/rockim/api/admin/tenant/v1/types"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/access/admin/module/tenant/biz"
	"rockimserver/app/access/admin/module/tenant/biz/options"
	"rockimserver/app/access/admin/module/tenant/service/converter"
)

type ProductService struct {
	uc        *biz.ProductUseCase
	sessionUc *biz.SessionUseCase
}

func NewProductService(uc *biz.ProductUseCase, sessionUc *biz.SessionUseCase) *ProductService {
	return &ProductService{uc: uc, sessionUc: sessionUc}
}

func (s *ProductService) Create(ctx context.Context, in *v1.ProductCreateRequest) (out *v1.ProductCreateResponse, err error) {
	session, err := s.sessionUc.Check(ctx)
	if err != nil {
		return
	}
	req := &options.ProductCreateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	req.TenantId = session.TenantId

	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}

	out = &v1.ProductCreateResponse{}
	return
}

func (s *ProductService) Update(ctx context.Context, in *v1.ProductUpdateRequest) (out *v1.ProductUpdateResponse, err error) {
	req := &options.ProductUpdateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	out = &v1.ProductUpdateResponse{}
	return
}

func (s *ProductService) Paging(ctx context.Context, in *v1.ProductPagingRequest) (out *v1.ProductPagingResponse, err error) {
	session, err := s.sessionUc.Check(ctx)
	if err != nil {
		return
	}
	req := &options.ProductPagingOptions{}
	req.TenantId = session.TenantId

	ret, err := s.uc.ListByTenant(ctx, req)
	if err != nil {
		return
	}
	list := make([]*types.Product, len(ret))
	for i, item := range ret {
		list[i] = converter.ToTenantProduct(item)
	}
	return &v1.ProductPagingResponse{List: list, Paginate: &shared.Paginated{Total: int64(len(list))}}, nil
}
