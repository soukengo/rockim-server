package entity

const (
	MongoFieldId = "_id"
)
const (
	TableSysTenantResource = "sys_tenant_resource"
)

type Table interface {
	TableName() string
}
