package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockim/api/logic/user/v1/types"
	"rockim/app/logic/user/data/database/convert"
	"rockim/app/logic/user/data/database/entity"
	"rockim/pkg/component/database/mongo"
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

func (d *UserData) FindByID(ctx context.Context, appId string, id string) (user *types.User, err error) {
	var record entity.ImUser
	err = d.mgo.FindOne(ctx, entity.TableImUser, bson.M{"appId": appId, entity.MongoFieldId: id}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.UserProto(&record), nil
}
func (d *UserData) FindByAccount(ctx context.Context, appId string, account string) (id string, err error) {
	projection := bson.M{entity.MongoFieldId: 1}
	var record entity.ImUser
	err = d.mgo.FindOne(ctx, entity.TableImUser, bson.M{"app_id": appId, "account": account}, &record, options.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id.Hex(), nil
}

func (d *UserData) Create(ctx context.Context, user *types.User) (err error) {
	record, err := convert.UserEntity(user)
	if err != nil {
		return
	}
	c := d.collection()
	_, err = c.InsertOne(ctx, record)
	return
}

func (d *UserData) Update(ctx context.Context, user *types.User) (err error) {
	record, err := convert.UserEntity(user)
	if err != nil {
		return
	}
	c := d.collection()
	_, err = c.UpdateByID(ctx, record.Id, record)
	return
}
