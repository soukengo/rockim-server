package options

import (
	"rockimserver/apis/rockim/shared/enums"
)

type SysTenantResourceOptions struct {
	Category enums.Admin_ResourceCategory
	// 菜单名称
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

type SysTenantResourceCreateOptions struct {
	Options *SysTenantResourceOptions
}

type SysTenantResourceUpdateOptions struct {
	Id      string
	Options *SysTenantResourceOptions
}
