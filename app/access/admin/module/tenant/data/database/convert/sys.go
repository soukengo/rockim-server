package convert

import (
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/access/admin/module/tenant/biz/types"
	"rockimserver/app/access/admin/module/tenant/data/database/entity"
)

func SysTenantResourceProto(source *entity.SysTenantResource) *types.SysTenantResource {
	return &types.SysTenantResource{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Category:   enums.Admin_ResourceCategory(source.Category),
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
