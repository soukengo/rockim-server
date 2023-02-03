package converter

import (
	apiTypes "rockimserver/apis/rockim/api/admin/tenant/v1/types"
	"rockimserver/apis/rockim/service/platform/v1/types"
)

func ToTenantProduct(source *types.Product) *apiTypes.Product {
	return &apiTypes.Product{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		TenantId:   source.TenantId,
		Name:       source.Name,
		ProductKey: source.ProductKey,
	}
}
