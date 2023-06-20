package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/apis/rockim/shared/enums"
	"rockimserver/app/logic/group/data/database/convert"
	"rockimserver/app/logic/group/data/database/entity"
)

type GroupData struct {
	mgo *mongo.Client
}

func NewChatRoomData(mgr *database.Manager) *GroupData {
	return &GroupData{mgo: mgr.Mongo()}
}
func (d *GroupData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableImGroup)
}

func (d *GroupData) FindByID(ctx context.Context, productId string, id string) (user *types.Group, err error) {
	var record entity.ImGroup
	err = d.mgo.FindOne(ctx, entity.TableImGroup, bson.M{"product_id": productId, mongo.FieldId: id}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.GroupProto(&record), nil
}

func (d *GroupData) FindGroupId(ctx context.Context, productId string, bizId string) (groupId string, err error) {
	projection := bson.M{mongo.FieldId: 1}
	var record entity.ImGroup
	err = d.mgo.FindOne(ctx, entity.TableImGroup, bson.M{"product_id": productId, "biz_id": bizId}, &record, options.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id, nil
}

func (d *GroupData) Create(ctx context.Context, user *types.Group) (err error) {
	record := convert.GroupEntity(user)
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *GroupData) Update(ctx context.Context, user *types.Group) (err error) {
	record := convert.GroupEntity(user)
	_, err = d.collection().UpdateByID(ctx, record.Id, record)
	return
}
func (d *GroupData) Delete(ctx context.Context, productId string, groupId string) (err error) {
	update := bson.M{"status": enums.Group_DELETED}
	_, err = d.collection().UpdateByID(ctx, groupId, update)
	return
}
