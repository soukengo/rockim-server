package entity

const (
	MongoFieldId = "_id"
)
const (
	TableTenant  = "tenant"
	TableProduct = "product"
)

type Table interface {
	TableName() string
}
