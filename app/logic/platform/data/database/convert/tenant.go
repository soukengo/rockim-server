package convert

import (
	"rockim/api/rockim/service/platform/v1/types"
	v1enums "rockim/api/rockim/shared/enums/v1"
	"rockim/app/logic/platform/data/database/entity"
)

func TenantProto(source *entity.Tenant) *types.Tenant {
	return &types.Tenant{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Name:       source.Name,
		Email:      source.Email,
		Password:   source.Password,
		Status:     v1enums.TenantStatus(source.Status),
	}
}
func TenantEntity(source *types.Tenant) *entity.Tenant {
	return &entity.Tenant{
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Name:       source.Name,
		Email:      source.Email,
		Password:   source.Password,
		Status:     int32(source.Status),
	}
}

func TenantResourceProto(source *entity.TenantResource) *types.TenantResource {
	return &types.TenantResource{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Category:   v1enums.TenantResourceCategory(source.Category),
		ParentId:   source.ParentId,
		Name:       source.Name,
		Url:        source.Url,
		Icon:       source.Icon,
		Order:      source.Order,
	}
}
func TenantResourceEntity(source *types.TenantResource) *entity.TenantResource {
	return &entity.TenantResource{
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		ParentId:   source.ParentId,
		Category:   int32(source.Category),
		Name:       source.Name,
		Url:        source.Url,
		Icon:       source.Icon,
		Order:      source.Order,
	}
}
