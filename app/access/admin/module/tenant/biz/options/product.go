package options

// ProductCreateOptions 创建请求
type ProductCreateOptions struct {
	// 所属租户
	TenantId string
	// 应用名称
	Name string
}

// ProductUpdateOptions 更新请求
type ProductUpdateOptions struct {
	// product id
	Id string
	// 应用名称
	Name string
}

type ProductPagingOptions struct {
	TenantId string
}
