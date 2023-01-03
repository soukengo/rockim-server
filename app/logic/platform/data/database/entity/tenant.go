package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tenant struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	Name       string             `bson:"name,omitempty"`        // 名称
	Email      string             `bson:"email,omitempty"`       // 邮箱
	Password   string             `bson:"password,omitempty"`    // 密码
	Status     int32              `bson:"status,omitempty"`      // 状态
}

func (*Tenant) TableName() string {
	return TableTenant
}

type TenantResource struct {
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

func (*TenantResource) TableName() string {
	return TablePlatResource
}
