package convert

import (
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/app/logic/platform/data/database/entity"
)

func ProductProto(source *entity.Product) *types.Product {
	return &types.Product{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		TenantId:   source.TenantId,
		Name:       source.Name,
		ProductKey: source.ProductKey,
	}
}
func ProductEntity(source *types.Product) *entity.Product {
	return &entity.Product{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		UpdateTime: source.UpdateTime,
		TenantId:   source.TenantId,
		Name:       source.Name,
		ProductKey: source.ProductKey,
	}
}
