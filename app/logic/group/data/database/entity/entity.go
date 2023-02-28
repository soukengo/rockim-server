package entity

const (
	TableImGroup = "im_group"
)

type Table interface {
	TableName() string
}
