package convert

import (
	"rockim/api/rockim/service/user/v1/types"
	v1enums "rockim/api/rockim/shared/enums/v1"
	"rockim/app/logic/user/data/database/entity"
)

func UserProto(source *entity.ImUser) *types.User {
	return &types.User{
		Id:         source.Id.Hex(),
		CreateTime: source.CreateTime,
		AppId:      source.ProductId,
		Bucket:     source.Bucket,
		Account:    source.Account,
		Name:       source.Name,
		Fields:     source.Fields,
		Status:     v1enums.UserStatus(source.Status),
	}
}
func UserEntity(source *types.User) *entity.ImUser {
	return &entity.ImUser{
		CreateTime: source.CreateTime,
		ProductId:  source.AppId,
		Bucket:     source.Bucket,
		Account:    source.Account,
		Name:       source.Name,
		Fields:     source.Fields,
		Status:     int32(source.Status),
	}
}
