package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/apis/rockim/service/user/v1/types"
	"rockimserver/app/logic/user/data/database/convert"
	"rockimserver/app/logic/user/data/database/entity"
	"rockimserver/pkg/component/database/mongo"
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

func (d *UserData) GenID(ctx context.Context) (string, error) {
	//seq, err := genMongoSeq(ctx, d.mgo, entity.TableImUser, true, 1)
	//if err != nil {
	//	return "", err
	//}
	//return strconv.FormatInt(seq, 10), nil
	return primitive.NewObjectID().Hex(), nil
}

func (d *UserData) FindByID(ctx context.Context, productId string, id string) (user *types.User, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.ImUser
	err = d.mgo.FindOne(ctx, entity.TableImUser, bson.M{"product_id": productId, entity.MongoFieldId: objId}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.UserProto(&record), nil
}
func (d *UserData) FindUidByAccount(ctx context.Context, productId string, account string) (id string, err error) {
	projection := bson.M{entity.MongoFieldId: 1}
	var record entity.ImUser
	err = d.mgo.FindOne(ctx, entity.TableImUser, bson.M{"product_id": productId, "account": account}, &record, options.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id.Hex(), nil
}

func (d *UserData) Create(ctx context.Context, user *types.User) (err error) {
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

func (d *UserData) Update(ctx context.Context, user *types.User) (err error) {
	objId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return
	}
	record := convert.UserEntity(user)
	record.Id = objId
	_, err = d.collection().UpdateByID(ctx, record.Id, record)
	return
}
