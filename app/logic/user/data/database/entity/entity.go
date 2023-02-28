package entity

const (
	TableImUser = "im_user"
)

type Table interface {
	TableName() string
}
