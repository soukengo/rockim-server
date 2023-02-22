package convert

import (
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"rockimserver/app/access/admin/module/manager/data/database/entity"
)

func UserProto(source *entity.SysUser) *types.SysUser {
	return &types.SysUser{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Account:    source.Account,
		Password:   source.Password,
		Name:       source.Name,
		IsAdmin:    source.IsAdmin,
	}
}
func UserEntity(source *types.SysUser) *entity.SysUser {
	return &entity.SysUser{
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Account:    source.Account,
		Password:   source.Password,
		Name:       source.Name,
		IsAdmin:    source.IsAdmin,
	}
}
func ResourceProto(source *entity.SysResource) *types.SysResource {
	return &types.SysResource{
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
func ResourceEntity(source *types.SysResource) *entity.SysResource {
	return &entity.SysResource{
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

func RoleProto(source *entity.SysRole) *types.SysRole {
	return &types.SysRole{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Name:       source.Name,
	}
}
func RoleEntity(source *types.SysRole) *entity.SysRole {
	return &entity.SysRole{
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Name:       source.Name,
	}
}

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
