package entity

const (
	MongoFieldId  = "_id"
	MongoFieldSeq = "seq"
)
const (
	TableImSeq  = "im_seq"
	TableImUser = "im_user"
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
