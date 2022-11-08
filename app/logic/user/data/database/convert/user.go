package convert

import (
	"rockim/api/logic/user/v1/types"
	"rockim/app/logic/user/data/database/entity"
)

func UserProto(source *entity.ImUser) *types.User {
	return &types.User{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		AppId:      source.AppId,
		Bucket:     source.Bucket,
		Account:    source.Account,
		Name:       source.Name,
		Fields:     source.Fields,
		Status:     types.UserStatus(source.Status),
	}
}
func UserEntity(source *types.User) *entity.ImUser {
	return &entity.ImUser{
		CreateTime: source.CreateTime,
		AppId:      source.AppId,
		Bucket:     source.Bucket,
		Account:    source.Account,
		Name:       source.Name,
		Fields:     source.Fields,
		Status:     int32(source.Status),
	}
}
