package biz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rockim/api/rockim/service/platform/v1/types"
	sharedTypes "rockim/api/rockim/shared/types"
	"rockim/pkg/util/copier"
	"time"
)

// PlatRoleRepo is a PlatRole repo.
type PlatRoleRepo interface {
	Create(ctx context.Context, record *types.PlatRole) error
	Update(ctx context.Context, record *types.PlatRole) error
	Delete(ctx context.Context, id string) error

	FindById(context.Context, string) (*types.PlatRole, error)
	ListByIds(ctx context.Context, ids []string) ([]*types.PlatRole, error)
	Paging(ctx context.Context, req *PlatRolePagingRequest) (res *PlatRolePagingResponse, err error)

	ListResourceId(ctx context.Context, ids []string) ([]string, error)
}

// PlatRoleUseCase is a PlatRole use case.
type PlatRoleUseCase struct {
	repo PlatRoleRepo
}

type PlatRoleOptions struct {
	Name string `json:"name,omitempty"` // 菜单名称
}

type PlatRoleCreateRequest struct {
	Options *PlatRoleOptions
}
type PlatRoleUpdateRequest struct {
	Id      string
	Options *PlatRoleOptions
}
type PlatRoleDeleteRequest struct {
	Id string
}

type PlatRolePagingRequest struct {
	Paginate *sharedTypes.Paginating
	Keyword  string
}
type PlatRolePagingResponse struct {
	List     []*types.PlatRole
	Paginate *sharedTypes.Paginated
}

// NewPlatRoleUseCase new a PlatRole use case.
func NewPlatRoleUseCase(repo PlatRoleRepo) *PlatRoleUseCase {
	return &PlatRoleUseCase{repo: repo}
}

func (uc *PlatRoleUseCase) Create(ctx context.Context, req *PlatRoleCreateRequest) (err error) {
	record := &types.PlatRole{Name: req.Options.Name}
	record.Id = primitive.NewObjectID().Hex()
	record.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, record)
}
func (uc *PlatRoleUseCase) Update(ctx context.Context, req *PlatRoleUpdateRequest) (err error) {
	record, err := uc.repo.FindById(ctx, req.Id)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(record, req.Options, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return
	}
	record.UpdateTime = time.Now().UnixMilli()
	return uc.repo.Update(ctx, record)
}
func (uc *PlatRoleUseCase) Delete(ctx context.Context, req *PlatRoleDeleteRequest) (err error) {
	return uc.repo.Delete(ctx, req.Id)
}
func (uc *PlatRoleUseCase) Paging(ctx context.Context, req *PlatRolePagingRequest) (res *PlatRolePagingResponse, err error) {
	return uc.repo.Paging(ctx, req)
}

func (uc *PlatRoleUseCase) ListByIds(ctx context.Context, ids []string) (res []*types.PlatRole, err error) {
	return uc.repo.ListByIds(ctx, ids)
}
func (uc *PlatRoleUseCase) ListResourceId(ctx context.Context, ids []string) ([]string, error) {
	return uc.repo.ListResourceId(ctx, ids)
}
