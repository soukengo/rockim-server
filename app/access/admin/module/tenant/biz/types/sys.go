package types

import (
	"rockimserver/apis/rockim/shared/enums"
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
	Category enums.Admin_ResourceCategory
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
