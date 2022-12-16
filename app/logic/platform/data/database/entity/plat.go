package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlatUser struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	Account    string             `bson:"account,omitempty"`     // 登录名
	Password   string             `bson:"password,omitempty"`    // 密码
	Name       string             `bson:"name,omitempty"`        // 用户名称
}

func (*PlatUser) TableName() string {
	return TablePlatUser
}

type PlatRole struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	Name       string             `bson:"name,omitempty"`        // 名称
}

func (*PlatRole) TableName() string {
	return TablePlatRole
}

type PlatResource struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	ParentId   string             `bson:"parent_id,omitempty"`   // 上级资源ID
	Name       string             `bson:"name,omitempty"`        // 资源名称
	Icon       string             `bson:"icon,omitempty"`        // ICON
	Url        string             `bson:"url,omitempty"`         // 资源URL
	Category   int32              `bson:"category,omitempty"`    // 类型 1：菜单，2：页面，3：功能
	Leaf       bool               `bson:"leaf,omitempty"`        // 是否叶子节点
	Order      int32              `bson:"order,omitempty"`       // 序号
	Level      int32              `bson:"level,omitempty"`       // 层级
}

func (*PlatResource) TableName() string {
	return TablePlatResource
}

type PlatUserRole struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UserId     string             `bson:"user_id,omitempty"`     // 用户ID
	RoleId     string             `bson:"role_id,omitempty"`     // 角色ID
}

func (*PlatUserRole) TableName() string {
	return TablePlatUserRole
}

type PlatRoleResource struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	RoleId     string             `bson:"role_id,omitempty"`     // 角色ID
	ResourceId string             `bson:"resource_id,omitempty"` // 资源ID
}

func (*PlatRoleResource) TableName() string {
	return TablePlatRoleResource
}
