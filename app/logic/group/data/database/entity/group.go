package entity

type ImGroup struct {
	Id            string            `bson:"_id"`                // ID
	CreateTime    int64             `bson:"create_time"`        // 创建时间
	UpdateTime    int64             `bson:"update_time"`        // 更新时间
	Category      int32             `bson:"category"`           // 类型
	ProductId     string            `bson:"product_id"`         // 所属应用
	CustomGroupId string            `bson:"custom_group_id"`    // 自定义群组ID
	Name          string            `bson:"name"`               // 名称
	IconUrl       string            `bson:"icon_url,omitempty"` // icon地址
	Fields        map[string]string `bson:"fields,omitempty"`   // 自定义字段
	Status        int32             `bson:"status"`             // 状态
	Owner         string            `bson:"owner,omitempty"`    // 群主
}

func (*ImGroup) TableName() string {
	return TableImGroup
}
