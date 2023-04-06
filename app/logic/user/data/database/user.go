package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/logic/user/data/database/convert"
	"rockimserver/app/logic/user/data/database/entity"
)

type UserData struct {
	mgo *mongo.Client
}

func NewUserData(mgo *mongo.Client) *UserData {
	return &UserData{mgo: mgo}
}

func (d *UserData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableImUser)
}

func (d *UserData) FindByID(ctx context.Context, productId string, id string) (user *types.User, err error) {
	var record entity.ImUser
	err = d.mgo.FindOne(ctx, entity.TableImUser, bson.M{"product_id": productId, mongo.FieldId: id}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.UserProto(&record), nil
}
func (d *UserData) FindUidByAccount(ctx context.Context, productId string, account string) (id string, err error) {
	projection := bson.M{mongo.FieldId: 1}
	var record entity.ImUser
	err = d.mgo.FindOne(ctx, entity.TableImUser, bson.M{"product_id": productId, "account": account}, &record, options.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id, nil
}

func (d *UserData) Create(ctx context.Context, user *types.User) (err error) {
	record := convert.UserEntity(user)
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *UserData) Update(ctx context.Context, user *types.User) (err error) {
	record := convert.UserEntity(user)
	_, err = d.collection().UpdateByID(ctx, record.Id, record)
	return
}
