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
