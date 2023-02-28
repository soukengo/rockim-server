package entity

const (
	TableImUser = "im_group"
)

type Table interface {
	TableName() string
}
