package options

import (
	"rockimserver/apis/rockim/shared"
)

// SysRoleOptions 角色选项
type SysRoleOptions struct {
	// 名称
	Name string
}

// SysRoleCreateOptions 角色创建请求
type SysRoleCreateOptions struct {
	Options *SysRoleOptions
}

// SysRoleUpdateOptions 角色更新请求
type SysRoleUpdateOptions struct {
	Id      string
	Options *SysRoleOptions
}

// SysRolePagingOptions 角色分页请求选项
type SysRolePagingOptions struct {
	Paginate *shared.Paginating
	Keyword  string
}
