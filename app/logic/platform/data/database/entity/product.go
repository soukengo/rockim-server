package entity

type Product struct {
	// 用户ID
	Id string `bson:"_id,omitempty"`
	// 创建时间
	CreateTime int64 `bson:"create_time,omitempty"`
	// 更新时间
	UpdateTime int64 `bson:"update_time,omitempty"`
	// 商户ID
	TenantId string `bson:"tenant_id,omitempty"`
	// 应用名称
	Name string `bson:"name,omitempty"`
	// 应用密钥
	ProductKey string `bson:"product_key,omitempty"`
}

func (*Product) TableName() string {
	return TableProduct
}
