package types

import (
	"rockim/api/rockim/shared/enums"
)

// SysUser 系统用户
type SysUser struct {
	// 系统用户ID
	Id string
	// 创建时间
	CreateTime int64
	// 创建时间
	UpdateTime int64
	// 登录名
	Account string
	// 密码
	Password string
	// 用户名称
	Name string
	// 是否超级管理员
	IsAdmin bool
}

// SysRole 系统角色
type SysRole struct {
	// 角色ID
	Id string
	// 创建时间
	CreateTime int64
	// 创建时间
	UpdateTime int64
	// 角色名称
	Name string
}

// SysResource 系统资源
type SysResource struct {
	// 资源ID
	Id string
	// 创建时间
	CreateTime int64
	// 创建时间
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
