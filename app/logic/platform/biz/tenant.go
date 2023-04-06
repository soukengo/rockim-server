package biz

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"github.com/soukengo/gopkg/log"
	"github.com/soukengo/gopkg/util/encrypt"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/apis/rockim/shared"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/logic/platform/biz/options"
	"time"
)

var (
	ErrTenantNotFound         = errors.NotFound(reasons.Platform_TENANT_NOT_FOUND.String(), "商户不存在")
	ErrTenantAccountDuplicate = errors.BadRequest(reasons.Platform_TENANT_ACCOUNT_DUPLICATE.String(), "不能重复创建")
)

type TenantRepo interface {
	GenID(ctx context.Context) (string, error)
	Create(ctx context.Context, record *types.Tenant) error
	Update(ctx context.Context, record *types.Tenant) error
	Delete(ctx context.Context, id string) error

	FindById(context.Context, string) (*types.Tenant, error)
	FindIdByEmail(ctx context.Context, account string) (id string, err error)
	Paging(ctx context.Context, opts *options.TenantPagingOptions) (ret []*types.Tenant, paginated *shared.Paginated, err error)
}

// TenantUseCase is a tenant use case.
type TenantUseCase struct {
	repo TenantRepo
}

func NewTenantUseCase(repo TenantRepo) *TenantUseCase {
	return &TenantUseCase{repo: repo}
}

func (uc *TenantUseCase) Create(ctx context.Context, opts *options.TenantCreateOptions) (err error) {
	oldId, err := uc.repo.FindIdByEmail(ctx, opts.Email)
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
	record := &types.Tenant{Name: opts.Name, Email: opts.Email, Password: opts.Password}
	plainPwd := opts.Password
	record.Id = newId
	record.Password = encrypt.MD5(plainPwd)
	record.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, record)
}
func (uc *TenantUseCase) Update(ctx context.Context, opts *options.TenantUpdateOptions) (err error) {
	record, err := uc.repo.FindById(ctx, opts.Id)
	if err != nil {
		log.WithContext(ctx).Errorf("FindById error: %v", err)
		err = ErrTenantNotFound
		return
	}
	record.Name = opts.Name
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

func (uc *TenantUseCase) Paging(ctx context.Context, opts *options.TenantPagingOptions) (ret []*types.Tenant, paginated *shared.Paginated, err error) {
	return uc.repo.Paging(ctx, opts)
}
