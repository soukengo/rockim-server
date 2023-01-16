package entity

type Tenant struct {
	Id         string `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64  `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64  `bson:"update_time,omitempty"` // 更新时间
	Name       string `bson:"name,omitempty"`        // 名称
	Email      string `bson:"email,omitempty"`       // 邮箱
	Password   string `bson:"password,omitempty"`    // 密码
	Status     int32  `bson:"status,omitempty"`      // 状态
}

func (*Tenant) TableName() string {
	return TableTenant
}
