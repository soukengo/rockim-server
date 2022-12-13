package convert

import (
	"rockim/api/rockim/service/platform/v1/types"
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
		AvatarUrl:  source.AvatarUrl,
	}
}
func UserEntity(source *types.PlatUser) *entity.PlatUser {
	return &entity.PlatUser{
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Account:    source.Account,
		Password:   source.Password,
		Name:       source.Name,
		AvatarUrl:  source.AvatarUrl,
	}
}
func ResourceProto(source *entity.PlatResource) *types.PlatResource {
	return &types.PlatResource{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		ParentId:   source.ParentId,
		Name:       source.Name,
		Url:        source.Url,
		Icon:       source.Icon,
		Level:      source.Level,
		Leaf:       source.Leaf,
		Order:      source.Order,
	}
}
func ResourceEntity(source *types.PlatResource) *entity.PlatResource {
	return &entity.PlatResource{
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		ParentId:   source.ParentId,
		Name:       source.Name,
		Url:        source.Url,
		Icon:       source.Icon,
		Level:      source.Level,
		Leaf:       source.Leaf,
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
