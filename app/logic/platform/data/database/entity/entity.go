package entity

const (
	MongoFieldId = "_id"
)
const (
	TablePlatUser         = "plat_user"
	TablePlatRole         = "plat_role"
	TablePlatResource     = "plat_resource"
	TablePlatUserRole     = "plat_user_role"
	TablePlatRoleResource = "plat_role_resource"

	TableTenant = "tenant"

	TableTenantResource = "tenant_resource"
)

type Table interface {
	TableName() string
}
