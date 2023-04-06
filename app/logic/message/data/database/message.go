package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/apis/rockim/service/message/v1/types"
	"rockimserver/app/logic/message/biz"
	biztypes "rockimserver/app/logic/message/biz/types"
	"rockimserver/app/logic/message/data/database/convert"
	"rockimserver/app/logic/message/data/database/entity"
	"time"
)

type MessageData struct {
	mgo *mongo.Client
}

func NewMessageSequenceData(mgo *mongo.Client) *MessageData {
	return &MessageData{mgo: mgo}
}

func (d *MessageData) GenSequence(ctx context.Context, productId string, conversationId *types.ConversationID) (int64, error) {
	id := biztypes.EncodeConversationID(productId, conversationId)
	seq, err := d.genMessageSequence(ctx, id, true)
	if err != nil {
		if !mgo.IsDuplicateKeyError(err) {
			return 0, err
		}
		// 如果报key重复错误，重新获取一次
		seq, err = d.genMessageSequence(ctx, id, false)
	}
	if seq == 0 {
		return 0, biz.ErrGenerateMessageFailed
	}
	return seq, nil
}

func (d *MessageData) genMessageSequence(ctx context.Context, id string, upsert bool) (seq int64, err error) {
	num := int64(1)
	returnDoc := options.After
	sr := d.mgo.Collection(entity.TableImMessageSequence).FindOneAndUpdate(ctx,
		bson.M{mongo.FieldId: id},
		bson.M{"$inc": bson.M{"seq": num}},
		&options.FindOneAndUpdateOptions{Upsert: &upsert, ReturnDocument: &returnDoc})
	var item entity.IMMessageSequence
	err = sr.Decode(&item)
	seq = item.Seq
	return
}

func (d *MessageData) Save(ctx context.Context, message *types.IMMessage) (err error) {
	record := convert.MessageEntity(message)
	record.UpdateTime = time.Now().UnixMilli()
	if err != nil {
		return
	}
	_, err = d.mgo.Collection(entity.TableImMessage).InsertOne(ctx, record)
	return
}

func (d *MessageData) SaveRelations(ctx context.Context, productId string, relations []*types.IMMessageRelation) (err error) {
	var records = make([]any, len(relations))
	for i, relation := range relations {
		records[i] = convert.MessageRelationEntity(relation)
	}
	_, err = d.mgo.Collection(entity.TableImMessageRelation).InsertMany(ctx, records)
	return
}
