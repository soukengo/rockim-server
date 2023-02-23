package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/app/logic/platform/data/database/entity"
	"rockimserver/pkg/component/database/mongo"
	"rockimserver/pkg/errors"
)

func genMongoSeq(ctx context.Context, client *mongo.Client, idType string, allowCreate bool, incr ...int64) (int64, error) {
	num := int64(1)
	if len(incr) > 0 {
		num = incr[0]
	}
	upsert := allowCreate
	returnDoc := options.After
	sr := client.Collection(entity.TableImSeq).FindOneAndUpdate(ctx,
		bson.M{entity.MongoFieldId: idType},
		bson.M{"$inc": bson.M{entity.MongoFieldSeq: num}},
		&options.FindOneAndUpdateOptions{Upsert: &upsert, ReturnDocument: &returnDoc})
	var item entity.ImSeq
	err := sr.Decode(&item)
	if err != nil {
		return 0, err
	}
	if item.Seq == 0 {
		return 0, errors.InternalServer("", "序号生成失败")
	}
	return item.Seq, nil
}
