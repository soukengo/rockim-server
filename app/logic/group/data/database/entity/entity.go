package entity

const (
	TableImGroup       = "im_group"
	TableImGroupMember = "im_group_member"
)

type Table interface {
	TableName() string
}
