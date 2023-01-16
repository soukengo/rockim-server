package types

import (
	"rockim/api/rockim/shared/enums"
)

// SysTenantResource 商户资源
type SysTenantResource struct {
	// 资源ID
	Id string
	// 创建时间
	CreateTime int64
	// 更新时间
	UpdateTime int64
	// 分类
	Category enums.AdminResourceCategory
	// 资源名称
	Name string
	// 上级ID
	ParentId string
	// URL
	Url string
	// ICON
	Icon string
	// 排序号
	Order int32
}
