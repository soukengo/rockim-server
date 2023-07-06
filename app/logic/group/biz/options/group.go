package options

import (
	"rockimserver/apis/rockim/shared"
	"rockimserver/apis/rockim/shared/enums"
)

type GroupMemberCheckOptions struct {
	ProductId string
	Category  enums.Group_Category
	GroupId   string
	Uid       string
}

type GroupMemberFindOptions struct {
	ProductId string
	Category  enums.Group_Category
	GroupId   string
	Uid       string
}
type GroupMemberListOptions struct {
	ProductId string
	Category  enums.Group_Category
	GroupId   string
	Uids      []string
}

type GroupMemberIdPaginateOptions struct {
	ProductId string
	Category  enums.Group_Category
	GroupId   string
	Paginate  *shared.Paginating
}
type GroupIdListByUidOptions struct {
	ProductId string
	Uid       string
}
