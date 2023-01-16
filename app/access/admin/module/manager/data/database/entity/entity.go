package entity

const (
	MongoFieldId = "_id"
)
const (
	TableSysUser           = "sys_user"
	TableSysRole           = "sys_role"
	TableSysResource       = "sys_resource"
	TableSysUserRole       = "sys_user_role"
	TableSysRoleResource   = "sys_role_resource"
	TableSysTenantResource = "sys_tenant_resource"
)

type Table interface {
	TableName() string
}
