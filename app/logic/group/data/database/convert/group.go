package convert

import (
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/group/data/database/entity"
)

func GroupProto(source *entity.ImGroup) *types.Group {
	return &types.Group{
		Id:            source.Id,
		CreateTime:    source.CreateTime,
		UpdateTime:    source.UpdateTime,
		Category:      enums.Group_Category(source.Category),
		ProductId:     source.ProductId,
		CustomGroupId: source.CustomGroupId,
		Name:          source.Name,
		IconUrl:       source.IconUrl,
		Fields:        source.Fields,
		Status:        enums.Group_Status(source.Status),
		Owner:         source.Owner,
	}
}
func GroupEntity(source *types.Group) *entity.ImGroup {
	return &entity.ImGroup{
		Id:            source.Id,
		CreateTime:    source.CreateTime,
		UpdateTime:    source.UpdateTime,
		Category:      int32(source.Category),
		ProductId:     source.ProductId,
		CustomGroupId: source.CustomGroupId,
		Name:          source.Name,
		IconUrl:       source.IconUrl,
		Fields:        source.Fields,
		Status:        int32(source.Status),
		Owner:         source.Owner,
	}
}
