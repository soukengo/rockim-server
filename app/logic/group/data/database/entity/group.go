package entity

type ImGroup struct {
	Id            string            `bson:"_id,omitempty"`             // ID
	CreateTime    int64             `bson:"create_time,omitempty"`     // 创建时间
	UpdateTime    int64             `bson:"update_time,omitempty"`     // 更新时间
	Category      int32             `bson:"category,omitempty"`        // 类型
	ProductId     string            `bson:"product_id,omitempty"`      // 所属应用
	CustomGroupId string            `bson:"custom_group_id,omitempty"` // 自定义群组ID
	Name          string            `bson:"name,omitempty"`            // 名称
	IconUrl       string            `bson:"icon_url,omitempty"`        // icon地址
	Fields        map[string]string `bson:"fields,omitempty"`          // 自定义字段
	Status        int32             `bson:"status,omitempty"`          // 状态
}

func (*ImGroup) TableName() string {
	return TableImGroup
}
