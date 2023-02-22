package options

import (
	"rockimserver/apis/rockim/shared/enums"
)

type SysResourceOptions struct {
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

type SysResourceCreateOptions struct {
	Options *SysResourceOptions
}

type SysResourceUpdateOptions struct {
	Id      string
	Options *SysResourceOptions
}
