package convert

import (
	"rockim/api/rockim/shared/enums"
	"rockim/app/access/admin/module/tenant/biz/types"
	"rockim/app/access/admin/module/tenant/data/database/entity"
)

func SysTenantResourceProto(source *entity.SysTenantResource) *types.SysTenantResource {
	return &types.SysTenantResource{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Category:   enums.AdminResourceCategory(source.Category),
		ParentId:   source.ParentId,
		Name:       source.Name,
		Url:        source.Url,
		Icon:       source.Icon,
		Order:      source.Order,
	}
}
func SysTenantResourceEntity(source *types.SysTenantResource) *entity.SysTenantResource {
	return &entity.SysTenantResource{
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
