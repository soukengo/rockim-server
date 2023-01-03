package biz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rockim/api/rockim/service/platform/v1/types"
	v1enums "rockim/api/rockim/shared/enums/v1"
	"rockim/pkg/util/copier"
	"time"
)

// TenantResourceRepo is a TenantResource repo.
type TenantResourceRepo interface {
	Create(context.Context, *types.TenantResource) error
	Update(context.Context, *types.TenantResource) error
	Delete(context.Context, string) error

	FindById(context.Context, string) (*types.TenantResource, error)
	List(context.Context) ([]*types.TenantResource, error)
	ListByIds(context.Context, []string) ([]*types.TenantResource, error)
}

// TenantResourceUseCase is a TenantResource use case.
type TenantResourceUseCase struct {
	repo TenantResourceRepo
}

// NewTenantResourceUseCase new a TenantResource use case.
func NewTenantResourceUseCase(repo TenantResourceRepo) *TenantResourceUseCase {
	return &TenantResourceUseCase{repo: repo}
}

type TenantResourceOptions struct {
	Category v1enums.TenantResourceCategory `json:"category,omitempty"`  // 类型
	Name     string                         `json:"name,omitempty"`      // 菜单名称
	ParentId string                         `json:"parent_id,omitempty"` // 上级ID
	Url      string                         `json:"url,omitempty"`       // URL
	Icon     string                         `json:"icon,omitempty"`      // ICON
	Level    int32                          `json:"level,omitempty"`     // 层级
	Leaf     bool                           `json:"leaf,omitempty"`      // 是否叶子节点
	Order    int32                          `json:"order,omitempty"`     // 排序号
}

type TenantResourceCreateRequest struct {
	Options *TenantResourceOptions
}
type TenantResourceUpdateRequest struct {
	Id      string
	Options *TenantResourceOptions
}
type TenantResourceDeleteRequest struct {
	Id string
}

func (uc *TenantResourceUseCase) Create(ctx context.Context, req *TenantResourceCreateRequest) (err error) {
	resource := &types.TenantResource{}
	err = copier.Copy(resource, req.Options)
	if err != nil {
		return
	}
	resource.Id = primitive.NewObjectID().Hex()
	resource.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, resource)
}
func (uc *TenantResourceUseCase) Update(ctx context.Context, req *TenantResourceUpdateRequest) (err error) {
	resource, err := uc.repo.FindById(ctx, req.Id)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(resource, req.Options, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return
	}
	resource.UpdateTime = time.Now().UnixMilli()
	return uc.repo.Update(ctx, resource)
}
func (uc *TenantResourceUseCase) Delete(ctx context.Context, req *TenantResourceDeleteRequest) (err error) {
	return uc.repo.Delete(ctx, req.Id)
}

func (uc *TenantResourceUseCase) List(ctx context.Context) (res []*types.TenantResource, err error) {
	return uc.repo.List(ctx)
}
func (uc *TenantResourceUseCase) ListByIds(ctx context.Context, ids []string) (res []*types.TenantResource, err error) {
	return uc.repo.ListByIds(ctx, ids)
}
