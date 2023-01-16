package biz

import (
	"context"
	v1 "rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/api/rockim/shared"
	"rockim/pkg/errors"
	"rockim/pkg/log"
	"rockim/pkg/util/encrypt"
	"time"
)

var (
	ErrTenantNotFound         = errors.NotFound(v1.ErrorReason_TENANT_NOT_FOUND.String(), "商户不存在")
	ErrTenantAccountDuplicate = errors.Conflict(v1.ErrorReason_TENANT_ACCOUNT_DUPLICATE.String(), "不能重复创建")
)

type TenantRepo interface {
	GenID(ctx context.Context) (string, error)
	Create(ctx context.Context, record *types.Tenant) error
	Update(ctx context.Context, record *types.Tenant) error
	Delete(ctx context.Context, id string) error

	FindById(context.Context, string) (*types.Tenant, error)
	FindIdByEmail(ctx context.Context, account string) (id string, err error)
	Paging(ctx context.Context, req *TenantPagingRequest) (res *TenantPagingResponse, err error)
}

type TenantCreateRequest struct {
	Email    string
	Password string
	Name     string
}
type TenantUpdateRequest struct {
	Id   string
	Name string
}

type TenantPagingRequest struct {
	Paginate *shared.Paginating
	Keyword  string
}
type TenantPagingResponse struct {
	List     []*types.Tenant
	Paginate *shared.Paginated
}

// TenantUseCase is a tenant use case.
type TenantUseCase struct {
	repo TenantRepo
}

func NewTenantUseCase(repo TenantRepo) *TenantUseCase {
	return &TenantUseCase{repo: repo}
}

func (uc *TenantUseCase) Create(ctx context.Context, req *TenantCreateRequest) (err error) {
	oldId, err := uc.repo.FindIdByEmail(ctx, req.Email)
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	if len(oldId) > 0 {
		return ErrTenantAccountDuplicate
	}
	newId, err := uc.repo.GenID(ctx)
	if err != nil {
		return
	}
	record := &types.Tenant{Name: req.Name, Email: req.Email, Password: req.Password}
	plainPwd := req.Password
	record.Id = newId
	record.Password = encrypt.MD5(plainPwd)
	record.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, record)
}
func (uc *TenantUseCase) Update(ctx context.Context, req *TenantUpdateRequest) (err error) {
	record, err := uc.repo.FindById(ctx, req.Id)
	if err != nil {
		log.WithContext(ctx).Errorf("FindById error: %v", err)
		err = ErrTenantNotFound
		return
	}
	record.Name = req.Name
	record.UpdateTime = time.Now().UnixMilli()
	return uc.repo.Update(ctx, record)
}

func (uc *TenantUseCase) Find(ctx context.Context, email string) (user *types.Tenant, err error) {
	uid, err := uc.repo.FindIdByEmail(ctx, email)
	if err != nil {
		log.WithContext(ctx).Errorf("FindIdByEmail email: %v error: %v", email, err)
		err = ErrTenantNotFound
		return
	}
	return uc.repo.FindById(ctx, uid)
}

func (uc *TenantUseCase) Paging(ctx context.Context, req *TenantPagingRequest) (res *TenantPagingResponse, err error) {
	return uc.repo.Paging(ctx, req)
}
