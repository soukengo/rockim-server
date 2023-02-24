package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ImUser struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`         // 用户ID
	CreateTime int64              `bson:"create_time,omitempty"` // 创建时间
	ProductId  string             `bson:"product_id,omitempty"`  // 所属应用
	Account    string             `bson:"account,omitempty"`     // 用户账户名，由接入方指定
	Name       string             `bson:"name,omitempty"`        // 用户名称
	AvatarUrl  string             `bson:"avatarUrl,omitempty"`   // 头像地址
	Status     int32              `bson:"status,omitempty"`      // 状态
	Fields     map[string]string  `bson:"fields,omitempty"`      // 客户自定义字段
}

func (*ImUser) TableName() string {
	return TableImUser
}
