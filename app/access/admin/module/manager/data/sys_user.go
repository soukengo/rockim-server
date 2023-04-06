package data

import (
	"context"
	"github.com/soukengo/gopkg/errors"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/access/admin/module/manager/biz"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"rockimserver/app/access/admin/module/manager/data/database"
)

type sysUserRepo struct {
	db *database.SysUserData
}

func NewSysUserRepo(db *database.SysUserData) biz.SysUserRepo {
	return &sysUserRepo{db: db}
}

func (r *sysUserRepo) Create(ctx context.Context, record *types.SysUser) (err error) {
	err = r.db.Create(ctx, record)
	return
}

func (r *sysUserRepo) Update(ctx context.Context, record *types.SysUser) (err error) {
	err = r.db.Update(ctx, record)
	return
}

func (r *sysUserRepo) Delete(ctx context.Context, id string) (err error) {
	err = r.db.Delete(ctx, id)
	return
}

func (r *sysUserRepo) FindById(ctx context.Context, id string) (*types.SysUser, error) {
	u, err := r.db.FindByID(ctx, id)
	if err != nil {
		if errors.IsNotFound(err) {
			err = biz.ErrSysUserNotFound
		}
	}
	return u, err
}
func (r *sysUserRepo) FindByAccount(ctx context.Context, account string) (*types.SysUser, error) {
	uid, err := r.db.FindIdByAccount(ctx, account)
	if err != nil {
		if errors.IsNotFound(err) {
			err = biz.ErrSysUserNotFound
		}
		return nil, err
	}
	return r.FindById(ctx, uid)
}

func (r *sysUserRepo) Paging(ctx context.Context, request *options.SysUserPagingOptions) (list []*types.SysUser, p *shared.Paginated, err error) {
	return r.db.Paging(ctx, request)
}

func (r *sysUserRepo) ListRoleId(ctx context.Context, userId string) ([]string, error) {
	return r.db.ListRoleId(ctx, userId)
}
func (r *sysUserRepo) SaveRoleId(ctx context.Context, userId string, roleIds []string) (err error) {
	err = r.db.SaveRoleId(ctx, userId, roleIds)
	return
}
