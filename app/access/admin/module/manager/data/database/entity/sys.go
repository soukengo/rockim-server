package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type SysUser struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	Account    string             `bson:"account,omitempty"`     // 登录名
	Password   string             `bson:"password,omitempty"`    // 密码
	Name       string             `bson:"name,omitempty"`        // 用户名称
	IsAdmin    bool               `bson:"is_admin,omitempty"`    // 是否超级管理员
}

func (*SysUser) TableName() string {
	return TableSysUser
}

type SysRole struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	Name       string             `bson:"name,omitempty"`        // 名称
}

func (*SysRole) TableName() string {
	return TableSysRole
}

type SysResource struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	ParentId   string             `bson:"parent_id,omitempty"`   // 上级资源ID
	Name       string             `bson:"name,omitempty"`        // 资源名称
	Icon       string             `bson:"icon,omitempty"`        // ICON
	Url        string             `bson:"url,omitempty"`         // 资源URL
	Category   int32              `bson:"category,omitempty"`    // 类型 1：菜单，2：页面，3：功能
	Order      int32              `bson:"order,omitempty"`       // 序号
}

func (*SysResource) TableName() string {
	return TableSysResource
}

type SysUserRole struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UserId     primitive.ObjectID `bson:"user_id,omitempty"`     // 用户ID
	RoleId     primitive.ObjectID `bson:"role_id,omitempty"`     // 角色ID
}

func (*SysUserRole) TableName() string {
	return TableSysUserRole
}

type SysRoleResource struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	RoleId     primitive.ObjectID `bson:"role_id,omitempty"`     // 角色ID
	ResourceId primitive.ObjectID `bson:"resource_id,omitempty"` // 资源ID
}

func (*SysRoleResource) TableName() string {
	return TableSysRoleResource
}

type SysTenantResource struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	ParentId   string             `bson:"parent_id,omitempty"`   // 上级资源ID
	Name       string             `bson:"name,omitempty"`        // 资源名称
	Icon       string             `bson:"icon,omitempty"`        // ICON
	Url        string             `bson:"url,omitempty"`         // 资源URL
	Category   int32              `bson:"category,omitempty"`    // 类型 1：菜单，2：页面，3：功能
	Order      int32              `bson:"order,omitempty"`       // 序号
}

func (*SysTenantResource) TableName() string {
	return TableSysTenantResource
}
