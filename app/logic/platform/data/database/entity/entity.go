package entity

const (
	MongoFieldId = "_id"
)
const (
	TablePlatUser = "plat_user"
)

type Table interface {
	TableName() string
}
