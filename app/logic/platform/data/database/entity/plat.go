package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlatUser struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	UpdateTime int64              `bson:"update_time,omitempty"` // 更新时间
	Account    string             `bson:"account,omitempty"`     // 登录名
	Password   string             `bson:"password,omitempty"`    // 密码
	Name       string             `bson:"name,omitempty"`        // 用户名称
	AvatarUrl  string             `bson:"avatarUrl,omitempty"`   // 头像
}

func (*PlatUser) TableName() string {
	return TablePlatUser
}
