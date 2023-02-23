package entity

const (
	MongoFieldId  = "_id"
	MongoFieldSeq = "seq"
)
const (
	TableImSeq   = "im_seq"
	TableTenant  = "tenant"
	TableProduct = "product"
)

type Table interface {
	TableName() string
}

type ImSeq struct {
	Id  string `bson:"_id,omitempty"` // Id
	Seq int64  `bson:"seq"`
}

func (*ImSeq) TableName() string {
	return TableImSeq
}
