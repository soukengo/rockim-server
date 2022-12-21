package convert

import (
	"rockim/api/rockim/service/platform/v1/types"
	v1enums "rockim/api/rockim/shared/enums/v1"
	"rockim/app/logic/platform/data/database/entity"
)

func UserProto(source *entity.PlatUser) *types.PlatUser {
	return &types.PlatUser{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Account:    source.Account,
		Password:   source.Password,
		Name:       source.Name,
		IsAdmin:    source.IsAdmin,
	}
}
func UserEntity(source *types.PlatUser) *entity.PlatUser {
	return &entity.PlatUser{
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Account:    source.Account,
		Password:   source.Password,
		Name:       source.Name,
		IsAdmin:    source.IsAdmin,
	}
}
func ResourceProto(source *entity.PlatResource) *types.PlatResource {
	return &types.PlatResource{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Category:   v1enums.PlatResourceCategory(source.Category),
		ParentId:   source.ParentId,
		Name:       source.Name,
		Url:        source.Url,
		Icon:       source.Icon,
		Order:      source.Order,
	}
}
func ResourceEntity(source *types.PlatResource) *entity.PlatResource {
	return &entity.PlatResource{
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

func RoleProto(source *entity.PlatRole) *types.PlatRole {
	return &types.PlatRole{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Name:       source.Name,
	}
}
func RoleEntity(source *types.PlatRole) *entity.PlatRole {
	return &entity.PlatRole{
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Name:       source.Name,
	}
}
