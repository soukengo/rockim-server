package convert

import (
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/platform/data/database/entity"
)

func TenantProto(source *entity.Tenant) *types.Tenant {
	return &types.Tenant{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Name:       source.Name,
		Email:      source.Email,
		Password:   source.Password,
		Status:     enums.TenantStatus(source.Status),
	}
}
func TenantEntity(source *types.Tenant) *entity.Tenant {
	return &entity.Tenant{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		Name:       source.Name,
		Email:      source.Email,
		Password:   source.Password,
		Status:     int32(source.Status),
	}
}
