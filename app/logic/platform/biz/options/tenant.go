package options

import "rockimserver/apis/rockim/shared"

type TenantCreateOptions struct {
	Email    string
	Password string
	Name     string
}

type TenantUpdateOptions struct {
	Id   string
	Name string
}

type TenantPagingOptions struct {
	Paginate *shared.Paginating
	Keyword  string
}
