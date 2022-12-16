package biz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rockim/api/rockim/service/platform/v1/types"
	v1enums "rockim/api/rockim/shared/enums/v1"
	"rockim/pkg/util/copier"
	"time"
)

// PlatResourceRepo is a PlatResource repo.
type PlatResourceRepo interface {
	Create(context.Context, *types.PlatResource) error
	Update(context.Context, *types.PlatResource) error
	Delete(context.Context, string) error

	FindById(context.Context, string) (*types.PlatResource, error)
	List(context.Context) ([]*types.PlatResource, error)
	ListByIds(context.Context, []string) ([]*types.PlatResource, error)
}

// PlatResourceUseCase is a PlatResource use case.
type PlatResourceUseCase struct {
	repo PlatResourceRepo
}

// NewPlatResourceUseCase new a PlatResource use case.
func NewPlatResourceUseCase(repo PlatResourceRepo) *PlatResourceUseCase {
	return &PlatResourceUseCase{repo: repo}
}

type PlatResourceOptions struct {
	Category v1enums.PlatResourceCategory `json:"category,omitempty"`  // 类型
	Name     string                       `json:"name,omitempty"`      // 菜单名称
	ParentId string                       `json:"parent_id,omitempty"` // 上级ID
	Url      string                       `json:"url,omitempty"`       // URL
	Icon     string                       `json:"icon,omitempty"`      // ICON
	Level    int32                        `json:"level,omitempty"`     // 层级
	Leaf     bool                         `json:"leaf,omitempty"`      // 是否叶子节点
	Order    int32                        `json:"order,omitempty"`     // 排序号
}

type PlatResourceCreateRequest struct {
	Options *PlatResourceOptions
}
type PlatResourceUpdateRequest struct {
	Id      string
	Options *PlatResourceOptions
}
type PlatResourceDeleteRequest struct {
	Id string
}

func (uc *PlatResourceUseCase) Create(ctx context.Context, req *PlatResourceCreateRequest) (err error) {
	resource := &types.PlatResource{}
	err = copier.Copy(resource, req.Options)
	if err != nil {
		return
	}
	resource.Id = primitive.NewObjectID().Hex()
	resource.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, resource)
}
func (uc *PlatResourceUseCase) Update(ctx context.Context, req *PlatResourceUpdateRequest) (err error) {
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
func (uc *PlatResourceUseCase) Delete(ctx context.Context, req *PlatResourceDeleteRequest) (err error) {
	return uc.repo.Delete(ctx, req.Id)
}

func (uc *PlatResourceUseCase) List(ctx context.Context) (res []*types.PlatResource, err error) {
	return uc.repo.List(ctx)
}
func (uc *PlatResourceUseCase) ListByIds(ctx context.Context, ids []string) (res []*types.PlatResource, err error) {
	return uc.repo.ListByIds(ctx, ids)
}
