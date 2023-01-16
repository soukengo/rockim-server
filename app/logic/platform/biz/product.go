package biz

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/pkg/errors"
	"rockim/pkg/log"
	"rockim/pkg/util/strings"
	"time"
)

var (
	ErrProductNotFound = errors.NotFound(v1.ErrorReason_APP_NOT_FOUND.String(), "应用不存在")
)

type ProductRepo interface {
	GenID(ctx context.Context) (string, error)
	Create(ctx context.Context, record *types.Product) error
	Update(ctx context.Context, record *types.Product) error

	FindById(context.Context, string) (*types.Product, error)
	ListByTenant(ctx context.Context, tenantId string) (res []*types.Product, err error)
}

type ProductCreateRequest struct {
	TenantId string
	Name     string
}
type ProductUpdateRequest struct {
	Id   string
	Name string
}

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(repo ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) Create(ctx context.Context, req *ProductCreateRequest) (err error) {
	id, err := uc.repo.GenID(ctx)
	if err != nil {
		return
	}
	record := &types.Product{TenantId: req.TenantId, Name: req.Name}
	record.Id = id
	record.ProductKey = strings.RandStr(32)
	record.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, record)
}

func (uc *ProductUseCase) Update(ctx context.Context, req *ProductUpdateRequest) (err error) {
	record, err := uc.repo.FindById(ctx, req.Id)
	if err != nil {
		log.WithContext(ctx).Errorf("FindById error: %v", err)
		err = ErrProductNotFound
		return
	}
	record.Name = req.Name
	record.UpdateTime = time.Now().UnixMilli()
	return uc.repo.Update(ctx, record)
}

func (uc *ProductUseCase) FindById(ctx context.Context, id string) (*types.Product, error) {
	return uc.repo.FindById(ctx, id)
}
func (uc *ProductUseCase) ListByTenant(ctx context.Context, tenantId string) (res []*types.Product, err error) {
	return uc.repo.ListByTenant(ctx, tenantId)
}
