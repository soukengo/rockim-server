package options

import (
	"rockimserver/apis/rockim/shared"
)

// SysUserOptions 用户选项
type SysUserOptions struct {
	// 名称
	Name string
}

// SysUserCreateOptions 用户创建请求
type SysUserCreateOptions struct {
	Options *SysUserOptions
	// 账号
	Account string
	// 密码
	Password string
}

// SysUserUpdateOptions 用户更新请求
type SysUserUpdateOptions struct {
	Id      string
	Options *SysUserOptions
}

// SysUserPagingOptions 用户分页请求
type SysUserPagingOptions struct {
	Paginate *shared.Paginating
	Keyword  string
}
