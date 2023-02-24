package convert

import (
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/user/data/database/entity"
)

func UserProto(source *entity.ImUser) *types.User {
	return &types.User{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		ProductId:  source.ProductId,
		Account:    source.Account,
		Name:       source.Name,
		AvatarUrl:  source.AvatarUrl,
		Fields:     source.Fields,
		Status:     enums.User_Status(source.Status),
	}
}
func UserEntity(source *types.User) *entity.ImUser {
	return &entity.ImUser{
		CreateTime: source.CreateTime,
		ProductId:  source.ProductId,
		Account:    source.Account,
		Name:       source.Name,
		AvatarUrl:  source.AvatarUrl,
		Fields:     source.Fields,
		Status:     int32(source.Status),
	}
}
