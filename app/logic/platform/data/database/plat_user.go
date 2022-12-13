package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/logic/platform/data/database/convert"
	"rockim/app/logic/platform/data/database/entity"
	"rockim/pkg/component/database/mongo"
)

type UserData struct {
	mgo *mongo.Client
}

func NewUserData(mgo *mongo.Client) *UserData {
	return &UserData{mgo: mgo}
}

func (d *UserData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TablePlatUser)
}

func (d *UserData) FindByID(ctx context.Context, id string) (user *types.PlatUser, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.PlatUser
	err = d.mgo.FindOne(ctx, entity.TablePlatUser, bson.M{entity.MongoFieldId: objId}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.UserProto(&record), nil
}

func (d *UserData) FindByAccount(ctx context.Context, account string) (id string, err error) {
	projection := bson.M{entity.MongoFieldId: 1}
	var record entity.PlatUser
	err = d.mgo.FindOne(ctx, entity.TablePlatUser, bson.M{"account": account}, &record, options.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id.Hex(), nil
}

func (d *UserData) Create(ctx context.Context, user *types.PlatUser) (err error) {
	objId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return
	}
	record := convert.UserEntity(user)
	record.Id = objId
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *UserData) Update(ctx context.Context, user *types.PlatUser) (err error) {
	objId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return
	}
	record := convert.UserEntity(user)
	record.Id = objId
	_, err = d.collection().UpdateByID(ctx, record.Id, record)
	return
}
