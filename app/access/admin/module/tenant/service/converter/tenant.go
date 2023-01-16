package converter

import (
	apiTypes "rockim/api/rockim/admin/tenant/v1/types"
	"rockim/api/rockim/service/platform/v1/types"
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
