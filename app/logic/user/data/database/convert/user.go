package convert

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func UserEntity(source *types.User) (*entity.ImUser, error) {
	id, err := primitive.ObjectIDFromHex(source.Id)
	if err != nil {
		return nil, err
	}
	return &entity.ImUser{
		Id:         id,
		CreateTime: source.CreateTime,
		AppId:      source.AppId,
		Bucket:     source.Bucket,
		Account:    source.Account,
		Name:       source.Name,
		Fields:     source.Fields,
		Status:     int32(source.Status),
	}, nil
}
