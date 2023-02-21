package biz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rockimserver/apis/rockim/shared"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"rockimserver/pkg/errors"
	"rockimserver/pkg/util/copier"
	"rockimserver/pkg/util/encrypt"
	"time"
)

const (
	defaultPassword = "123456"
)

var (
	ErrSysUserNotFound         = errors.NotFound(reasons.Admin_SYS_USER_NOT_FOUND.String(), "账号不存在")
	ErrSysUserAccountDuplicate = errors.BadRequest(reasons.Admin_SYS_USER_DUPLICATE.String(), "不能重复创建")
)

type SysUserRepo interface {
	Create(context.Context, *types.SysUser) error
	Update(context.Context, *types.SysUser) error
	Delete(context.Context, string) error

	FindById(ctx context.Context, id string) (*types.SysUser, error)
	FindByAccount(ctx context.Context, account string) (*types.SysUser, error)
	Paging(context.Context, *options.SysUserPagingOptions) ([]*types.SysUser, *shared.Paginated, error)

	ListRoleId(ctx context.Context, userId string) ([]string, error)
	SaveRoleId(ctx context.Context, userId string, roleIds []string) error
}

type SysUserUseCase struct {
	repo SysUserRepo
}

func NewSysUserUseCase(repo SysUserRepo) *SysUserUseCase {
	return &SysUserUseCase{repo: repo}
}

func (uc *SysUserUseCase) Create(ctx context.Context, opts *options.SysUserCreateOptions) (err error) {
	old, err := uc.repo.FindByAccount(ctx, opts.Account)
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	if old != nil {
		return ErrSysUserAccountDuplicate
	}
	err = nil
	record := &types.SysUser{Name: opts.Options.Name, Account: opts.Account, Password: opts.Password}
	plainPwd := opts.Password
	if len(plainPwd) == 0 {
		plainPwd = defaultPassword
	}
	record.Password = encrypt.MD5(plainPwd)
	record.Id = primitive.NewObjectID().Hex()
	record.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, record)
}
func (uc *SysUserUseCase) Update(ctx context.Context, opts *options.SysUserUpdateOptions) (err error) {
	record, err := uc.repo.FindById(ctx, opts.Id)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(record, opts.Options, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return
	}
	record.UpdateTime = time.Now().UnixMilli()
	return uc.repo.Update(ctx, record)
}
func (uc *SysUserUseCase) Delete(ctx context.Context, id string) (err error) {
	return uc.repo.Delete(ctx, id)
}
func (uc *SysUserUseCase) Paging(ctx context.Context, opts *options.SysUserPagingOptions) ([]*types.SysUser, *shared.Paginated, error) {
	return uc.repo.Paging(ctx, opts)
}
func (uc *SysUserUseCase) ListRoleId(ctx context.Context, userId string) ([]string, error) {
	return uc.repo.ListRoleId(ctx, userId)
}

func (uc *SysUserUseCase) SaveRoleId(ctx context.Context, userId string, roleIds []string) error {
	return uc.repo.SaveRoleId(ctx, userId, roleIds)
}
