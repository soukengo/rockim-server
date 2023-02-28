package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/app/logic/group/data/database/entity"
	"rockimserver/pkg/component/database/mongo"
)

type ChatRoomData struct {
	mgo *mongo.Client
}

func NewChatRoomData(mgo *mongo.Client) *ChatRoomData {
	return &ChatRoomData{mgo: mgo}
}

func (d *ChatRoomData) FindByCustomGroupId(ctx context.Context, productId string, customGroupId string) (groupId string, err error) {
	projection := bson.M{mongo.FieldId: 1}
	var record entity.ImGroup
	err = d.mgo.FindOne(ctx, entity.TableImUser, bson.M{"product_id": productId, "custom_group_id": customGroupId}, &record, options.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id, nil
}
